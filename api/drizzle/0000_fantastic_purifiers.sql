CREATE TABLE "user_times" (
	"id" integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY (sequence name "user_times_id_seq" INCREMENT BY 1 MINVALUE 1 MAXVALUE 2147483647 START WITH 1 CACHE 1),
	"user" text NOT NULL,
	"start_time" text NOT NULL,
	"end_time" text NOT NULL,
	"hours_spent" integer NOT NULL,
	"minutes_spent" integer NOT NULL,
	"seconds_spent" integer NOT NULL,
	"created_at" timestamp DEFAULT now()
);
