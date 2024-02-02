import {
  Column,
  CreateDateColumn,
  Entity,
  OneToMany,
  PrimaryGeneratedColumn,
} from 'typeorm';
import { OrderItem } from './order-item';

export enum OrderStatus {
  PENDING = 'pending',
  PAID = 'paid',
  FAILED = 'falied',
}

export type CreateOrderCommand = {
  // operacao de negocio, e nao salvar no banco
  client_id: number;
  items: {
    product_id: string;
    quantity: number;
    price: number;
  }[];
};

@Entity()
export class Order {
  @PrimaryGeneratedColumn('uuid')
  id: string;

  @Column()
  total: number;

  @Column()
  client_id: number;

  @Column()
  status: OrderStatus = OrderStatus.PENDING;

  @CreateDateColumn()
  created_at: Date;

  @OneToMany(() => OrderItem, (item) => item.order, {
    // ao fazer o relacionamento salva-se a ordem e os itens de uma vez
    cascade: ['insert'],
    eager: true,
  })
  items: OrderItem[];

  static create(input: CreateOrderCommand) {
    const order = new Order();
    order.client_id = input.client_id;
    order.items = input.items.map((item) => {
      const orderItem = new OrderItem();
      orderItem.product_id = item.product_id;
      orderItem.quatity = item.quantity;
      orderItem.price = item.price;
      return orderItem;
    });
    order.total = order.items.reduce((sum, item) => {
      return sum + item.price * item.quatity;
    }, 0);
    return order;
  }
}
