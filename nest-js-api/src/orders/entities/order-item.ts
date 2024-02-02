import { Product } from 'src/products/entities/product.entity';
import { Order } from './order.entity';
import {
  Column,
  Entity,
  JoinColumn,
  ManyToOne,
  PrimaryGeneratedColumn,
} from 'typeorm';

@Entity()
export class OrderItem {
  @PrimaryGeneratedColumn('uuid')
  id: number;
  @Column({ type: 'int' })
  quatity: number;
  @Column({ type: 'decimal', precision: 10, scale: 2 })
  price: number;
  @ManyToOne(() => Product)
  @JoinColumn({ name: 'product_id' })
  product: Product;
  @Column()
  product_id: string;
  @ManyToOne(() => Order)
  order: Order;
}
