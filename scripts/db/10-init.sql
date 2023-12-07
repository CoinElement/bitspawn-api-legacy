CREATE DATABASE bitspawn OWNER postgres;
\c bitspawn;
CREATE TABLE IF NOT EXISTS public.admin_configs
(
    rpc_node_url  text COLLATE pg_catalog."default",
    faucet_key    text COLLATE pg_catalog."default",
    spwn_contract text COLLATE pg_catalog."default",
    fee_contract  text COLLATE pg_catalog."default",
    wallet_url    text COLLATE pg_catalog."default",
    id            integer NOT NULL,
    created_at    timestamp with time zone,
    updated_at    timestamp with time zone,
    admin_user_id text COLLATE pg_catalog."default",
    adapter_url   text COLLATE pg_catalog."default",
    organizer_url text COLLATE pg_catalog."default",
    cred_contract text COLLATE pg_catalog."default",
    cred_fee      text COLLATE pg_catalog."default",
    spwn_fee      text COLLATE pg_catalog."default",
    api_rate      bigint,
    CONSTRAINT admin_configs_pkey PRIMARY KEY (id)
);

INSERT INTO admin_configs(rpc_node_url, faucet_key, spwn_contract, fee_contract, wallet_url, id, created_at, updated_at,
                          admin_user_id, adapter_url, organizer_url, cred_contract, cred_fee, spwn_fee, api_rate)
VALUES ('http://prod.elb.amazonaws.com', '1111111111111111111111111', '0x11111111111', '0xC1111111111',
        'https://dev-wallet.bitspawn.gg', 1, '2020-04-16 16:43:51.256677+00', '2021-05-04 20:01:11.287164+00',
        'larry@bitspawn.gg', 'http://111.amazonaws.com/v1/', 'https://dev-tournament-organizer.bitspawn.gg', '0x111111',
        '0x111111', '0x111111', null);


CREATE TABLE IF NOT EXISTS public.user_accounts
(
    username       text COLLATE pg_catalog."default" NOT NULL,
    created_at     timestamp with time zone,
    updated_at     timestamp with time zone,
    deleted_at     timestamp with time zone,
    public_address text COLLATE pg_catalog."default",
    display_name   text COLLATE pg_catalog."default",
    is_active      boolean,
    avatar_url     text COLLATE pg_catalog."default",
    sub            text COLLATE pg_catalog."default" NOT NULL,
    eth_gifted_at  timestamp with time zone,
    online_time    timestamp with time zone,
    favourite      jsonb,
    profile_banner text COLLATE pg_catalog."default",
    phone_number   text COLLATE pg_catalog."default",
    enabled2_fa    boolean,
    CONSTRAINT user_accounts_pkey PRIMARY KEY (sub),
    CONSTRAINT user_accounts_display_name_key UNIQUE (display_name)
);

INSERT INTO public.user_accounts(username, created_at, updated_at, deleted_at, public_address, display_name, is_active,
                                 avatar_url, sub, eth_gifted_at, online_time, favourite, profile_banner, phone_number,
                                 enabled2_fa)
VALUES ('test@bitspawn.gg', '2020-12-07 18:11:06.506106+00', '2020-12-07 18:11:06.506106+00', null,
        '0xf462227e87d91a4c1567ed3b8681e0da0ffd5a445ea992fcd6dcf2d2e982b41c', 'lionfish', true, '',
        'c025acff-5961-4973-ad3b-da803c828549', null, null, null, null, null, null);

INSERT INTO public.user_accounts(username, created_at, updated_at, deleted_at, public_address, display_name, is_active,
                                 avatar_url, sub, eth_gifted_at, online_time, favourite, profile_banner, phone_number,
                                 enabled2_fa)
VALUES ('test1@bitspawn.gg', '2020-12-07 18:11:06.506106+00', '2020-12-07 18:11:06.506106+00', null,
        '0xf462227e87d91a4c1567ed3b8681e0da0ffd5a445ea992fcd6dcf2d2e982b41c', 'jelyfish', true, '',
        '4a231fa0-3a08-4607-8b30-9d0f27160b86', null, null, null, null, null, null);
