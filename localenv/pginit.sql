create table asset (
                       id serial primary key,
                       symbol text not null,
                       name text not null,
                       address text,
                       decimals    text,
                    unique (symbol, address)

);