--
-- PostgreSQL database dump
--

-- Dumped from database version 16.10 (Debian 16.10-1.pgdg13+1)
-- Dumped by pg_dump version 17.0

-- Started on 2025-11-25 16:44:07

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

DROP DATABASE nextmed_db;
--
-- TOC entry 3443 (class 1262 OID 16384)
-- Name: nextmed_db; Type: DATABASE; Schema: -; Owner: nextmed_user
--

CREATE DATABASE nextmed_db WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.utf8';


ALTER DATABASE nextmed_db OWNER TO nextmed_user;

\connect nextmed_db

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 215 (class 1259 OID 16473)
-- Name: admins; Type: TABLE; Schema: public; Owner: nextmed_user
--

CREATE TABLE public.admins (
    id uuid NOT NULL,
    email character varying(255) NOT NULL,
    password text NOT NULL,
    full_name character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.admins OWNER TO nextmed_user;

--
-- TOC entry 216 (class 1259 OID 16484)
-- Name: clients; Type: TABLE; Schema: public; Owner: nextmed_user
--

CREATE TABLE public.clients (
    client_id uuid NOT NULL,
    name character varying(100) NOT NULL,
    email character varying(255) NOT NULL,
    api_key character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.clients OWNER TO nextmed_user;

--
-- TOC entry 217 (class 1259 OID 16568)
-- Name: logs; Type: TABLE; Schema: public; Owner: nextmed_user
--

CREATE TABLE public.logs (
    log_id uuid NOT NULL,
    client_id uuid NOT NULL,
    api_key character varying(255) NOT NULL,
    ip character varying(45),
    endpoint text NOT NULL,
    "timestamp" timestamp without time zone DEFAULT now() NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.logs OWNER TO nextmed_user;

--
-- TOC entry 3435 (class 0 OID 16473)
-- Dependencies: 215
-- Data for Name: admins; Type: TABLE DATA; Schema: public; Owner: nextmed_user
--

INSERT INTO public.admins VALUES ('acd1f353-5622-48f4-9fd8-031af7244d90', 'admin@example.com', '$2b$12$JK8pvJIwa8/VrJtAfTuMTOvyjOpcfWPWLNTV9FzQP.vkluLiZaAe2', 'Admin', '2025-11-25 06:32:17.559719', '2025-11-25 06:32:17.559719', NULL);


--
-- TOC entry 3436 (class 0 OID 16484)
-- Dependencies: 216
-- Data for Name: clients; Type: TABLE DATA; Schema: public; Owner: nextmed_user
--

INSERT INTO public.clients VALUES ('0991fd21-91cb-497d-9a45-13573300b026', 'rizky@nex.com', 'rizky', '088b01a996499e307f064e393496e62569e40e53c6f61319fe67553f4be90067', '2025-11-25 08:45:50.984736', '2025-11-25 08:45:50.984736', NULL);


--
-- TOC entry 3437 (class 0 OID 16568)
-- Dependencies: 217
-- Data for Name: logs; Type: TABLE DATA; Schema: public; Owner: nextmed_user
--

INSERT INTO public.logs VALUES ('4d467d2d-7a16-482c-af54-b16c9a9be7cb', '0991fd21-91cb-497d-9a45-13573300b026', '088b01a996499e307f064e393496e62569e40e53c6f61319fe67553f4be90067', '192.168.1.22', '/api/v1/client/register', '2025-11-25 08:51:06.344067', '2025-11-25 08:51:06.344067', '2025-11-25 08:51:06.344067', NULL);
INSERT INTO public.logs VALUES ('51f70aa2-0642-41a8-8361-54666895c836', '0991fd21-91cb-497d-9a45-13573300b026', '088b01a996499e307f064e393496e62569e40e53c6f61319fe67553f4be90067', '192.168.1.22', '/api/v1/client/register', '2025-11-25 08:51:07.364238', '2025-11-25 08:51:07.364238', '2025-11-25 08:51:07.364238', NULL);
INSERT INTO public.logs VALUES ('380a61cc-1bab-4aa8-8e29-ef316edd9b8b', '0991fd21-91cb-497d-9a45-13573300b026', '088b01a996499e307f064e393496e62569e40e53c6f61319fe67553f4be90067', '192.168.1.22', '/api/v1/client/register', '2025-11-25 08:51:07.956124', '2025-11-25 08:51:07.956124', '2025-11-25 08:51:07.956124', NULL);
INSERT INTO public.logs VALUES ('06f3b15a-451a-484b-a25e-4716f1096432', '0991fd21-91cb-497d-9a45-13573300b026', '088b01a996499e307f064e393496e62569e40e53c6f61319fe67553f4be90067', '192.168.1.22', '/api/v1/client/register', '2025-11-25 08:51:08.392519', '2025-11-25 08:51:08.392519', '2025-11-25 08:51:08.392519', NULL);
INSERT INTO public.logs VALUES ('212f818b-eeec-4f97-bbe4-c33aba472dad', '0991fd21-91cb-497d-9a45-13573300b026', '088b01a996499e307f064e393496e62569e40e53c6f61319fe67553f4be90067', '192.168.1.22', '/api/v1/client/register', '2025-11-25 08:51:09.078507', '2025-11-25 08:51:09.078507', '2025-11-25 08:51:09.078507', NULL);
INSERT INTO public.logs VALUES ('b2ede6cd-93c7-4186-892c-781777bca4ba', '0991fd21-91cb-497d-9a45-13573300b026', '088b01a996499e307f064e393496e62569e40e53c6f61319fe67553f4be90067', '192.168.1.22', '/api/v1/client/register', '2025-11-25 08:51:09.520763', '2025-11-25 08:51:09.520763', '2025-11-25 08:51:09.520763', NULL);
INSERT INTO public.logs VALUES ('a0ea0e79-1a3c-4984-80a8-918a9952825b', '0991fd21-91cb-497d-9a45-13573300b026', '088b01a996499e307f064e393496e62569e40e53c6f61319fe67553f4be90067', '192.168.1.22', '/api/v1/client/register', '2025-11-25 14:54:19.841672', '2025-11-25 14:54:19.841672', '2025-11-25 14:54:19.841672', NULL);
INSERT INTO public.logs VALUES ('1bc4816f-6ce1-4d4e-bdc8-f594883cdd99', '0991fd21-91cb-497d-9a45-13573300b026', '088b01a996499e307f064e393496e62569e40e53c6f61319fe67553f4be90067', '192.168.1.22', '/api/v1/client/register', '2025-11-25 14:54:21.227099', '2025-11-25 14:54:21.227099', '2025-11-25 14:54:21.227099', NULL);
INSERT INTO public.logs VALUES ('69ffaf87-02ad-4a33-8a48-dba3d4e6938a', '0991fd21-91cb-497d-9a45-13573300b026', '088b01a996499e307f064e393496e62569e40e53c6f61319fe67553f4be90067', '', '/api/v1/client/register', '2025-11-25 14:54:34.099866', '2025-11-25 14:54:34.099866', '2025-11-25 14:54:34.099866', NULL);
INSERT INTO public.logs VALUES ('512dd374-6b84-4cc5-9392-32c670d3166b', '0991fd21-91cb-497d-9a45-13573300b026', '088b01a996499e307f064e393496e62569e40e53c6f61319fe67553f4be90067', '192.168.1.22', '/api/v1/client/register', '2025-11-25 14:55:53.109675', '2025-11-25 14:55:53.109675', '2025-11-25 14:55:53.109675', NULL);
INSERT INTO public.logs VALUES ('339285ca-be3b-4c01-a978-6ba97eb0b895', '0991fd21-91cb-497d-9a45-13573300b026', '088b01a996499e307f064e393496e62569e40e53c6f61319fe67553f4be90067', '192.168.1.22', '/api/v1/client/register', '2025-11-25 15:23:21.787294', '2025-11-25 15:23:21.787294', '2025-11-25 15:23:21.787294', NULL);
INSERT INTO public.logs VALUES ('9869e712-6b89-4ad9-993b-ac8582df0e13', '0991fd21-91cb-497d-9a45-13573300b026', '088b01a996499e307f064e393496e62569e40e53c6f61319fe67553f4be90067', '192.168.1.22', '/api/v1/client/register', '2025-11-25 15:23:22.881142', '2025-11-25 15:23:22.881142', '2025-11-25 15:23:22.881142', NULL);
INSERT INTO public.logs VALUES ('122f2348-ce2d-46a9-85fb-b62ba42aff59', '0991fd21-91cb-497d-9a45-13573300b026', '088b01a996499e307f064e393496e62569e40e53c6f61319fe67553f4be90067', '192.168.1.22', '/api/v1/client/register', '2025-11-25 15:23:23.685512', '2025-11-25 15:23:23.685512', '2025-11-25 15:23:23.685512', NULL);
INSERT INTO public.logs VALUES ('e4369067-760c-4f3c-b73d-807644ab19ce', '0991fd21-91cb-497d-9a45-13573300b026', '088b01a996499e307f064e393496e62569e40e53c6f61319fe67553f4be90067', '192.168.1.22', '/api/v1/client/register', '2025-11-25 15:23:25.591601', '2025-11-25 15:23:25.591601', '2025-11-25 15:23:25.591601', NULL);
INSERT INTO public.logs VALUES ('d3861d2d-2971-47e0-8351-31d6d2e62f25', '0991fd21-91cb-497d-9a45-13573300b026', '088b01a996499e307f064e393496e62569e40e53c6f61319fe67553f4be90067', '192.168.1.22', '/api/v1/client/register', '2025-11-25 15:23:26.214685', '2025-11-25 15:23:26.214685', '2025-11-25 15:23:26.214685', NULL);
INSERT INTO public.logs VALUES ('9684b920-3b90-4848-a178-2ede1a86ab11', '0991fd21-91cb-497d-9a45-13573300b026', '088b01a996499e307f064e393496e62569e40e53c6f61319fe67553f4be90067', '192.168.1.22', '/api/v1/client/register', '2025-11-25 16:12:55.611716', '2025-11-25 16:12:55.611716', '2025-11-25 16:12:55.611716', NULL);
INSERT INTO public.logs VALUES ('ff2b4146-8df7-45e9-a88b-2f9913e146b2', '0991fd21-91cb-497d-9a45-13573300b026', '088b01a996499e307f064e393496e62569e40e53c6f61319fe67553f4be90067', '192.168.1.22', '/api/v1/client/register', '2025-11-25 16:12:56.326375', '2025-11-25 16:12:56.326375', '2025-11-25 16:12:56.326375', NULL);
INSERT INTO public.logs VALUES ('3fef60d9-f39f-4ee3-9213-7553811af6b8', '0991fd21-91cb-497d-9a45-13573300b026', '088b01a996499e307f064e393496e62569e40e53c6f61319fe67553f4be90067', '192.168.1.22', '/api/v1/client/register', '2025-11-25 16:12:56.903929', '2025-11-25 16:12:56.903929', '2025-11-25 16:12:56.903929', NULL);
INSERT INTO public.logs VALUES ('93b32fad-9104-47fb-a8c8-3cd9c5668dc6', '0991fd21-91cb-497d-9a45-13573300b026', '088b01a996499e307f064e393496e62569e40e53c6f61319fe67553f4be90067', '192.168.1.22', '/api/v1/client/register', '2025-11-25 16:12:57.432767', '2025-11-25 16:12:57.432767', '2025-11-25 16:12:57.432767', NULL);
INSERT INTO public.logs VALUES ('de9e7ebb-a606-4e24-8836-d326be65e2d2', '0991fd21-91cb-497d-9a45-13573300b026', '088b01a996499e307f064e393496e62569e40e53c6f61319fe67553f4be90067', '192.168.1.22', '/api/v1/client/register', '2025-11-25 16:12:57.91678', '2025-11-25 16:12:57.91678', '2025-11-25 16:12:57.91678', NULL);
INSERT INTO public.logs VALUES ('5f0300ce-dbf5-4b1a-8ac2-a2a26f51893e', '0991fd21-91cb-497d-9a45-13573300b026', '088b01a996499e307f064e393496e62569e40e53c6f61319fe67553f4be90067', '192.168.1.22', '/api/v1/client/register', '2025-11-25 16:13:33.925216', '2025-11-25 16:13:33.925216', '2025-11-25 16:13:33.925216', NULL);
INSERT INTO public.logs VALUES ('b4bf4090-df7d-4515-843c-72ba9a18feaf', '0991fd21-91cb-497d-9a45-13573300b026', '088b01a996499e307f064e393496e62569e40e53c6f61319fe67553f4be90067', '192.168.1.22', '/api/v1/client/register', '2025-11-25 16:13:40.096652', '2025-11-25 16:13:40.096652', '2025-11-25 16:13:40.096652', NULL);


--
-- TOC entry 3282 (class 2606 OID 16483)
-- Name: admins admins_email_key; Type: CONSTRAINT; Schema: public; Owner: nextmed_user
--

ALTER TABLE ONLY public.admins
    ADD CONSTRAINT admins_email_key UNIQUE (email);


--
-- TOC entry 3284 (class 2606 OID 16481)
-- Name: admins admins_pkey; Type: CONSTRAINT; Schema: public; Owner: nextmed_user
--

ALTER TABLE ONLY public.admins
    ADD CONSTRAINT admins_pkey PRIMARY KEY (id);


--
-- TOC entry 3286 (class 2606 OID 16494)
-- Name: clients clients_email_key; Type: CONSTRAINT; Schema: public; Owner: nextmed_user
--

ALTER TABLE ONLY public.clients
    ADD CONSTRAINT clients_email_key UNIQUE (email);


--
-- TOC entry 3288 (class 2606 OID 16492)
-- Name: clients clients_pkey; Type: CONSTRAINT; Schema: public; Owner: nextmed_user
--

ALTER TABLE ONLY public.clients
    ADD CONSTRAINT clients_pkey PRIMARY KEY (client_id);


--
-- TOC entry 3290 (class 2606 OID 16577)
-- Name: logs logs_pkey; Type: CONSTRAINT; Schema: public; Owner: nextmed_user
--

ALTER TABLE ONLY public.logs
    ADD CONSTRAINT logs_pkey PRIMARY KEY (log_id);


--
-- TOC entry 3291 (class 2606 OID 16578)
-- Name: logs logs_client_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nextmed_user
--

ALTER TABLE ONLY public.logs
    ADD CONSTRAINT logs_client_id_fkey FOREIGN KEY (client_id) REFERENCES public.clients(client_id);


-- Completed on 2025-11-25 16:44:07

--
-- PostgreSQL database dump complete
--

