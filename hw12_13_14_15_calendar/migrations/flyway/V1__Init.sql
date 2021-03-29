create table events (
                        id text primary key,
                        title text,
                        description text,
                        authorId text,
                        start_date date,
                        start_time time,
                        end_date date,
                        end_time time,
                        notification_time time
);
create index authorId_idx on events (authorId);
create index start_idx on events using btree (start_date, start_time);