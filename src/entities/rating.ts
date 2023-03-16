import {
  Entity,
  Column,
  PrimaryColumn,
  CreateDateColumn,
  UpdateDateColumn,
  Double,
} from 'typeorm'

@Entity()
export default class Ratings {
  @PrimaryColumn()
  ratingId: string

  @Column()
  userId: string

  @Column()
  movieId: string

  @Column({ type: 'float4' })
  ratingValue: number

  @CreateDateColumn({
    type: 'timestamp',
    default: () => 'CURRENT_TIMESTAMP(6)',
  })
  public created_at: Date

  @UpdateDateColumn({
    type: 'timestamp',
    default: () => 'CURRENT_TIMESTAMP(6)',
    onUpdate: 'CURRENT_TIMESTAMP(6)',
  })
  public updated_at: Date
}
