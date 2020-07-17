--
-- PostgreSQL database dump
--

-- Dumped from database version 12.3
-- Dumped by pg_dump version 12.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: db_covid_palembang; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA db_covid_palembang;


ALTER SCHEMA db_covid_palembang OWNER TO postgres;

--
-- Name: adminpack; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS adminpack WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION adminpack; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION adminpack IS 'administrative functions for PostgreSQL';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: daerah_kecamatan; Type: TABLE; Schema: db_covid_palembang; Owner: postgres
--

CREATE TABLE db_covid_palembang.daerah_kecamatan (
    id integer NOT NULL,
    kecamatan character varying(100),
    jumlah_penduduk_positif integer,
    jumlah_penduduk_pulih integer,
    jumlah_penduduk_wafat integer,
    keadaan_zona integer
);


ALTER TABLE db_covid_palembang.daerah_kecamatan OWNER TO postgres;

--
-- Name: daerah_kecamatan_id_seq; Type: SEQUENCE; Schema: db_covid_palembang; Owner: postgres
--

CREATE SEQUENCE db_covid_palembang.daerah_kecamatan_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE db_covid_palembang.daerah_kecamatan_id_seq OWNER TO postgres;

--
-- Name: daerah_kecamatan_id_seq; Type: SEQUENCE OWNED BY; Schema: db_covid_palembang; Owner: postgres
--

ALTER SEQUENCE db_covid_palembang.daerah_kecamatan_id_seq OWNED BY db_covid_palembang.daerah_kecamatan.id;


--
-- Name: zona_daerah; Type: TABLE; Schema: db_covid_palembang; Owner: postgres
--

CREATE TABLE db_covid_palembang.zona_daerah (
    id integer NOT NULL,
    zona_daerah character varying(100)
);


ALTER TABLE db_covid_palembang.zona_daerah OWNER TO postgres;

--
-- Name: zona_daerah_id_seq; Type: SEQUENCE; Schema: db_covid_palembang; Owner: postgres
--

CREATE SEQUENCE db_covid_palembang.zona_daerah_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE db_covid_palembang.zona_daerah_id_seq OWNER TO postgres;

--
-- Name: zona_daerah_id_seq; Type: SEQUENCE OWNED BY; Schema: db_covid_palembang; Owner: postgres
--

ALTER SEQUENCE db_covid_palembang.zona_daerah_id_seq OWNED BY db_covid_palembang.zona_daerah.id;


--
-- Name: daerah_kecamatan id; Type: DEFAULT; Schema: db_covid_palembang; Owner: postgres
--

ALTER TABLE ONLY db_covid_palembang.daerah_kecamatan ALTER COLUMN id SET DEFAULT nextval('db_covid_palembang.daerah_kecamatan_id_seq'::regclass);


--
-- Name: zona_daerah id; Type: DEFAULT; Schema: db_covid_palembang; Owner: postgres
--

ALTER TABLE ONLY db_covid_palembang.zona_daerah ALTER COLUMN id SET DEFAULT nextval('db_covid_palembang.zona_daerah_id_seq'::regclass);


--
-- Data for Name: daerah_kecamatan; Type: TABLE DATA; Schema: db_covid_palembang; Owner: postgres
--

COPY db_covid_palembang.daerah_kecamatan (id, kecamatan, jumlah_penduduk_positif, jumlah_penduduk_pulih, jumlah_penduduk_wafat, keadaan_zona) FROM stdin;
17	5 ULU	10	12	0	2
108	ALANG ALANG LEBAR	10	0	0	2
56	18 ILIR	0	12	0	1
57	16 ILIR	0	12	0	1
62	KEPANDEAN BARU	0	12	0	1
63	20 ILIR I	0	12	0	1
60	15 ILIR	0	12	0	1
61	17 ILIR	0	12	0	1
66	20 ILIR IV	0	12	0	1
67	SEKIP JAYA	0	12	0	1
64	SEI PANGERAN	0	12	0	1
65	20 ILIR III	0	12	0	1
95	SRIMULYO	0	0	0	1
82	KALIDONI	10	0	0	2
71	TALANG AMAN	10	12	0	2
70	PIPA REJA	10	12	0	2
100	SUKARAMI	10	0	0	2
89	8 ILIR	0	12	0	1
88	9 ILIR	0	12	0	1
90	SUKAMAJU	0	12	0	1
75	1 ILIR	10	10	0	2
84	10 ILIR	0	12	0	1
73	LAWANG KIDUL	10	10	0	2
3	32 ILIR	0	10	0	1
85	11 ILIR	0	12	0	1
76	SUNGAI BUAH	10	10	0	2
87	DUKU	0	12	0	1
4	30 ILIR	0	10	0	1
86	KUTO BATU	0	12	1	1
23	OGAN BARU	10	0	0	2
80	SEI SELINCAH	10	0	0	2
26	TUAN KENTANG	10	0	0	2
27	8 ULU	10	0	0	2
24	KERTAPATI	10	0	0	2
25	15 ULU	10	0	0	2
78	5 ILIR	10	0	0	2
28	SILABERANTI	10	0	0	2
97	KARYA MULYA	0	0	0	1
46	LOROK PAKJO	10	12	1	2
29	9/10 ULU	10	12	0	2
51	22 ILIR	10	12	1	2
94	LEBONG GAJAH	0	10	0	1
77	2 ILIR	10	0	0	2
18	7 ULU	10	12	0	2
47	DEMANG LEBAR DAUN	10	12	0	2
44	BUKIT LAMA	10	12	0	2
45	26 ILIR I	10	12	0	2
50	TALANG SEMUT	10	12	0	2
48	BUKIT BARU	10	12	0	2
49	SIRING AGUNG	10	12	0	2
92	SAKO	0	10	0	1
52	19 ILIR	10	12	0	2
107	TALANG KELAPA	10	0	0	2
106	KARYA BARU	0	0	0	1
104	TALANG JAMBE	0	0	0	1
105	SRIJAYA	0	0	0	1
102	TALANG BETUTU	0	10	0	1
103	SUKODADI	0	10	0	1
101	KEBUN BUNGA	0	10	0	1
53	23 ILIR	10	12	0	2
8	27 ILIR	0	0	1	1
30	SENTOSA	0	12	0	1
91	SIALANG	0	10	1	1
69	20 ILIR II	10	12	1	2
79	SEI LAIS	10	0	0	2
6	29 ILIR	0	0	0	1
7	28 ILIR	0	0	0	1
68	PAHLAWAN	10	12	0	2
5	KEMANG MANIS	0	0	0	1
10	GANDUS	0	0	0	1
13	36 ILIR	0	12	1	1
72	ARIO KEMUNING	10	12	0	2
9	PULO KERTO	0	0	0	1
81	SEI SELAYUR	10	0	0	2
83	BUKIT SANGKAL	10	12	0	2
96	SUKA MULYA	0	0	0	1
11	KARANG JAYA	0	12	0	1
98	SUKA BANGUN	10	0	0	2
99	SUKAJAYA	10	0	0	2
14	1 ULU	0	12	0	1
15	2 ULU	0	12	0	1
12	KARANG ANYAR	0	12	0	1
74	3 ILIR	10	10	1	2
16	3-4 ULU	0	12	0	1
93	SAKO BARU	0	10	0	1
21	KEMANG AGUNG	10	10	0	2
20	KERAMASAN	10	10	0	2
22	KEMAS RINDO	10	10	0	2
31	16 ULU	0	12	0	1
2	35 ILIR	10	10	0	2
19	KARYA JAYA	10	10	0	2
34	13 ULU	0	12	0	1
35	12 ULU	0	12	0	1
32	TANGGA TAKAT	0	12	0	1
33	14 ULU	0	12	0	1
38	TALANG PUTRI	0	12	0	1
39	KOMPERTA	0	12	0	1
36	11 ULU	0	12	0	1
37	PLAJU DARAT	0	12	0	1
42	PLAJU ULU	0	12	0	1
43	BAGUS KUNING	0	12	0	1
40	PLAJU ILIR	0	12	0	1
41	TALANG BUBUK	0	12	0	1
54	26 ILIR	0	12	0	1
55	24 ILIR	0	12	0	1
58	13 ILIR	0	12	0	1
59	14 ILIR	0	12	0	1
\.


--
-- Data for Name: zona_daerah; Type: TABLE DATA; Schema: db_covid_palembang; Owner: postgres
--

COPY db_covid_palembang.zona_daerah (id, zona_daerah) FROM stdin;
1	Zona Hijau
2	Zona Kuning
3	Zona Merah
4	Zona Hitam
\.


--
-- Name: daerah_kecamatan_id_seq; Type: SEQUENCE SET; Schema: db_covid_palembang; Owner: postgres
--

SELECT pg_catalog.setval('db_covid_palembang.daerah_kecamatan_id_seq', 108, true);


--
-- Name: zona_daerah_id_seq; Type: SEQUENCE SET; Schema: db_covid_palembang; Owner: postgres
--

SELECT pg_catalog.setval('db_covid_palembang.zona_daerah_id_seq', 4, true);


--
-- Name: daerah_kecamatan daerah_kecamatan_pk; Type: CONSTRAINT; Schema: db_covid_palembang; Owner: postgres
--

ALTER TABLE ONLY db_covid_palembang.daerah_kecamatan
    ADD CONSTRAINT daerah_kecamatan_pk PRIMARY KEY (id);


--
-- Name: zona_daerah zona_daerah_pk; Type: CONSTRAINT; Schema: db_covid_palembang; Owner: postgres
--

ALTER TABLE ONLY db_covid_palembang.zona_daerah
    ADD CONSTRAINT zona_daerah_pk PRIMARY KEY (id);


--
-- Name: daerah_kecamatan daerah_kecamatan_zona_daerah_id_fk; Type: FK CONSTRAINT; Schema: db_covid_palembang; Owner: postgres
--

ALTER TABLE ONLY db_covid_palembang.daerah_kecamatan
    ADD CONSTRAINT daerah_kecamatan_zona_daerah_id_fk FOREIGN KEY (keadaan_zona) REFERENCES db_covid_palembang.zona_daerah(id);


--
-- PostgreSQL database dump complete
--

