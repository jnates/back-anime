-- CREATE TABLE TYPE-IDENTIFIER
-- Table: public.type_identifiers

-- DROP TABLE public.type_identifiers;

CREATE TABLE public.types_identifiers
(
    type_id integer NOT NULL,
    type_description character varying COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT types_identifiers_pkey PRIMARY KEY (type_id)
)

TABLESPACE pg_default;

ALTER TABLE public.types_identifiers
    OWNER to postgres;

COMMENT ON COLUMN public.types_identifiers.type_id
    IS 'Tipo de identificación del personaje';

COMMENT ON COLUMN public.types_identifiers.type_description
    IS 'Descripción del tipo de identificador';

-- CREATE TABLE USERS
-- Table: public.users

-- DROP TABLE public.users;

CREATE TABLE public.users
(
    user_id integer NOT NULL,
    user_name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    user_identifier integer NOT NULL,
    user_email character varying(255) COLLATE pg_catalog."default" NOT NULL,
    user_password character varying(255) COLLATE pg_catalog."default" NOT NULL,
    user_type_identifier integer NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (user_id),
    CONSTRAINT fk_type_identifier FOREIGN KEY (user_type_identifier)
        REFERENCES public.types_identifiers (type_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE public.users
    OWNER to postgres;

COMMENT ON COLUMN public.users.user_id
    IS 'LLave primaria de Usuarios';

COMMENT ON COLUMN public.users.user_name
    IS 'Nombre completo del Personaje';

COMMENT ON COLUMN public.users.user_identifier
    IS 'Numero de identificación de usuario';

COMMENT ON COLUMN public.users.user_email
    IS 'Correo electronico del Usuario para el login';

COMMENT ON COLUMN public.users.user_password
    IS 'Password de usuario para login';

COMMENT ON COLUMN public.users.user_type_identifier
    IS 'Tipo de documento de identificacion del usuario';

-- CREATE TABLE PRODUCTS
-- -- Table: public.products

-- DROP TABLE public.products;

CREATE TABLE public.products
(
    product_id integer NOT NULL,
    product_name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    product_amount integer,
    product_user_created integer NOT NULL,
    product_date_created timestamp(0) without time zone NOT NULL,
    product_user_modify integer NOT NULL,
    "product_date_modify" timestamp(0) without time zone NOT NULL,
    CONSTRAINT product_pkey PRIMARY KEY (product_id),
    CONSTRAINT fk_user_created FOREIGN KEY (product_user_created)
        REFERENCES public.users (user_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT fk_user_modify FOREIGN KEY (product_user_modify)
        REFERENCES public.users (user_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)

TABLESPACE pg_default;

ALTER TABLE public.products
    OWNER to postgres;

COMMENT ON COLUMN public.products.product_id
    IS 'id del producto';

COMMENT ON COLUMN public.products.product_name
    IS 'Name del producto';

COMMENT ON COLUMN public.products.product_amount
    IS 'Cantidad del producto';

COMMENT ON COLUMN public.products.product_user_created
    IS 'El usuario crea un producto';

COMMENT ON COLUMN public.products.product_date_created
    IS 'Fecha de creación del producto';

COMMENT ON COLUMN public.products.product_date_modify
    IS 'Fecha de  modificación del producto';

COMMENT ON COLUMN public.products."product_user_modify"
    IS 'Fecha de modificación del user';