import { pgTable, text, integer, timestamp } from 'drizzle-orm/pg-core';

export const userTimes = pgTable('user_times', {
  id: integer('id').primaryKey().generatedAlwaysAsIdentity(),
  user: text('user').notNull(),
  startTime: text('start_time').notNull(),
  endTime: text('end_time').notNull(),
  hoursSpent: integer('hours_spent').notNull(),
  minutesSpent: integer('minutes_spent').notNull(),
  secondsSpent: integer('seconds_spent').notNull(),
  createdAt: timestamp('created_at').defaultNow()
});
