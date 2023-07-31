create table test
(
    id    uuid not null
        constraint test_pk
            primary key,
    title text not null
);

insert into postgres.public.test (id, title) values (gen_random_uuid(), '(2) Test entry 6');
insert into postgres.public.test (id, title) values (gen_random_uuid(), '(2) Test entry 7');
insert into postgres.public.test (id, title) values (gen_random_uuid(), '(2) Test entry 8');
insert into postgres.public.test (id, title) values (gen_random_uuid(), '(2) Test entry 9');
insert into postgres.public.test (id, title) values (gen_random_uuid(), '(2) Test entry 0');