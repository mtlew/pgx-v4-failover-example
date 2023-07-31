create table test
(
    id    uuid not null
        constraint test_pk
            primary key,
    title text not null
);

insert into postgres.public.test (id, title) values (gen_random_uuid(), '(1) Test entry 1');
insert into postgres.public.test (id, title) values (gen_random_uuid(), '(1) Test entry 2');
insert into postgres.public.test (id, title) values (gen_random_uuid(), '(1) Test entry 3');
insert into postgres.public.test (id, title) values (gen_random_uuid(), '(1) Test entry 4');
insert into postgres.public.test (id, title) values (gen_random_uuid(), '(1) Test entry 5');