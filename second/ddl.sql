-- Project Name : Personal-LINK
-- Date/Time    : 2023/08/04 13:37:40
-- Author       : sd
-- RDBMS Type   : PostgreSQL
-- Application  : A5:SQL Mk-2

-- 許諾管理テーブル
CREATE TABLE consent_table (
  id bigserial NOT NULL
  , consent_id text NOT NULL
  , consent_request_setting_id text NOT NULL
  , base_id text NOT NULL
  , consent_status text NOT NULL
  , consent_flag boolean
  , consent_request_date_time timestamp(6) with time zone NOT NULL
  , consent_end_date_time timestamp(6) with time zone
  , consent_date_time timestamp(6) with time zone
  , cancel_date_time timestamp(6) with time zone
  , denial_date_time timestamp(6) with time zone
  , CONSTRAINT consent_table_PKC PRIMARY KEY (id)
) ;

CREATE UNIQUE INDEX consent_table_IX1
  ON consent_table(consent_id);

-- RP連携サービス管理テーブル
CREATE TABLE rp_service_link_table (
  service_link_id bigserial NOT NULL
  , base_id text NOT NULL
  , linking_member_code character(7) NOT NULL
  , linking_subsystem_code text NOT NULL
  , id_service_code text NOT NULL
  , id_service_userid text NOT NULL
  , uax_link_connect_flag boolean
  , linking_date_time timestamp(6) with time zone NOT NULL
  , CONSTRAINT rp_service_link_table_PKC PRIMARY KEY (service_link_id)
) ;

CREATE UNIQUE INDEX rp_service_link_table_IX1
  ON rp_service_link_table(base_id,linking_member_code,linking_subsystem_code,id_service_code,id_service_userid);

-- 許諾依頼設定マスタ
CREATE TABLE consent_request_setting_master (
  consent_request_id bigserial NOT NULL
  , consent_request_setting_id text NOT NULL
  , open_flag boolean NOT NULL
  , consent_type text NOT NULL
  , consent_end_date_time timestamp(6) with time zone NOT NULL
  , data_provider_member_code character(7) NOT NULL
  , data_provider_subsystem_code text NOT NULL
  , data_provider_service_code text NOT NULL
  , service_provider_member_code character(7) NOT NULL
  , service_provider_subsystem_code text NOT NULL
  , purpose_of_use text NOT NULL
  , merit_of_consent text NOT NULL
  , create_date timestamp(6) with time zone NOT NULL
  , update_date timestamp(6) with time zone NOT NULL
  , CONSTRAINT consent_request_setting_master_PKC PRIMARY KEY (consent_request_id)
) ;

CREATE UNIQUE INDEX consent_request_setting_master_IX1
  ON consent_request_setting_master(consent_request_setting_id);

-- ID変換マスタ
CREATE TABLE id_convert_master (
  id_convert_id bigserial NOT NULL
  , base_id text NOT NULL
  , main_id_flag boolean NOT NULL
  , id_service_code text NOT NULL
  , id_service_userid text NOT NULL
  , email text NOT NULL
  , email_verified text NOT NULL
  , ial text NOT NULL
  , aal text NOT NULL
  , fal text NOT NULL
  , linking_date_time timestamp(6) with time zone NOT NULL
  , CONSTRAINT id_convert_master_PKC PRIMARY KEY (id_convert_id)
) ;

CREATE UNIQUE INDEX id_convert_master_IX1
  ON id_convert_master(base_id,id_service_userid,email);

-- IDサービスマスタ
CREATE TABLE id_service_master (
  id_service_id bigserial NOT NULL
  , id_service_code text NOT NULL
  , id_service_name text NOT NULL
  , id_service_logo oid
  , id_provider_name text NOT NULL
  , create_date_time timestamp(6) with time zone NOT NULL
  , update_date_time timestamp(6) with time zone
  , CONSTRAINT id_service_master_PKC PRIMARY KEY (id_service_id)
) ;

CREATE UNIQUE INDEX id_service_master_IX1
  ON id_service_master(id_service_code);

-- データサービスマスタ
CREATE TABLE service_master (
  service_id bigserial NOT NULL
  , member_code character(7) NOT NULL
  , subsystem_code text NOT NULL
  , service_code text NOT NULL
  , service_version text NOT NULL
  , service_name text NOT NULL
  , description text NOT NULL
  , input_parameters xml NOT NULL
  , output_parameters xml NOT NULL
  , create_date timestamp(6) with time zone NOT NULL
  , update_date timestamp(6) with time zone NOT NULL
  , CONSTRAINT service_master_PKC PRIMARY KEY (service_id)
) ;

CREATE UNIQUE INDEX service_master_IX1
  ON service_master(member_code,subsystem_code,service_code);

-- サブシステムマスタ
CREATE TABLE subsystem_master (
  subsystem_id bigserial NOT NULL
  , member_code character(7) NOT NULL
  , subsystem_code text NOT NULL
  , subsystem_name text NOT NULL
  , service_url text NOT NULL
  , service_logo_file oid NOT NULL
  , description text NOT NULL
  , create_date timestamp(6) with time zone NOT NULL
  , update_date timestamp(6) with time zone NOT NULL
  , CONSTRAINT subsystem_master_PKC PRIMARY KEY (subsystem_id)
) ;

CREATE UNIQUE INDEX subsystem_master_IX1
  ON subsystem_master(member_code,subsystem_code);

-- ユーザマスタ
CREATE TABLE user_master (
  id bigserial NOT NULL
  , base_id text NOT NULL
  , email text NOT NULL
  , main_id_service text NOT NULL
  , kyc_flag boolean NOT NULL
  , latest_ial text
  , latest_kyc_date_time timestamp(6) with time zone
  , mobile_phone text
  , mobile_phone_validation boolean
  , terms_of_service_consent_flag boolean
  , terms_of_service_consent_date_time timestamp(6) with time zone
  , create_date_time timestamp(6) with time zone NOT NULL
  , update_date_time timestamp(6) with time zone
  , CONSTRAINT user_master_PKC PRIMARY KEY (id)
) ;

CREATE UNIQUE INDEX user_master_IX1
  ON user_master(base_id);

-- メンバーコードマスタ
CREATE TABLE member_code_master (
  member_id bigserial NOT NULL
  , member_class character(3) NOT NULL
  , member_code character(7) NOT NULL
  , member_name text NOT NULL
  , description text NOT NULL
  , create_date timestamp(6) with time zone NOT NULL
  , update_date timestamp(6) with time zone NOT NULL
  , CONSTRAINT member_code_master_PKC PRIMARY KEY (member_id)
) ;

CREATE UNIQUE INDEX member_code_master_IX1
  ON member_code_master(member_code);

ALTER TABLE consent_request_setting_master
  ADD CONSTRAINT consent_request_setting_master_FK1 FOREIGN KEY (service_provider_member_code,service_provider_subsystem_code) REFERENCES subsystem_master(member_code,subsystem_code);

ALTER TABLE consent_request_setting_master
  ADD CONSTRAINT consent_request_setting_master_FK2 FOREIGN KEY (data_provider_member_code,data_provider_subsystem_code,data_provider_service_code) REFERENCES service_master(member_code,subsystem_code,service_code);

ALTER TABLE consent_table
  ADD CONSTRAINT consent_table_FK1 FOREIGN KEY (base_id) REFERENCES user_master(base_id);

ALTER TABLE consent_table
  ADD CONSTRAINT consent_table_FK2 FOREIGN KEY (consent_request_setting_id) REFERENCES consent_request_setting_master(consent_request_setting_id);

ALTER TABLE id_convert_master
  ADD CONSTRAINT id_convert_master_FK1 FOREIGN KEY (id_service_code) REFERENCES id_service_master(id_service_code);

ALTER TABLE id_convert_master
  ADD CONSTRAINT id_convert_master_FK2 FOREIGN KEY (base_id) REFERENCES user_master(base_id);

CREATE UNIQUE INDEX id_convert_master_IX2
  ON id_convert_master(base_id,id_service_code,id_service_userid);

ALTER TABLE rp_service_link_table
  ADD CONSTRAINT rp_service_link_table_FK1 FOREIGN KEY (base_id,id_service_code,id_service_userid) REFERENCES id_convert_master(base_id,id_service_code,id_service_userid);

ALTER TABLE rp_service_link_table
  ADD CONSTRAINT rp_service_link_table_FK2 FOREIGN KEY (linking_member_code,linking_subsystem_code) REFERENCES subsystem_master(member_code,subsystem_code);

ALTER TABLE rp_service_link_table
  ADD CONSTRAINT rp_service_link_table_FK3 FOREIGN KEY (base_id) REFERENCES user_master(base_id);

ALTER TABLE service_master
  ADD CONSTRAINT service_master_FK1 FOREIGN KEY (member_code,subsystem_code) REFERENCES subsystem_master(member_code,subsystem_code);

ALTER TABLE subsystem_master
  ADD CONSTRAINT subsystem_master_FK1 FOREIGN KEY (member_code) REFERENCES member_code_master(member_code);

COMMENT ON TABLE consent_table IS '許諾管理テーブル';
COMMENT ON COLUMN consent_table.id IS 'id';
COMMENT ON COLUMN consent_table.consent_id IS '許諾ID';
COMMENT ON COLUMN consent_table.consent_request_setting_id IS '許諾依頼設定ID';
COMMENT ON COLUMN consent_table.base_id IS 'BaseID';
COMMENT ON COLUMN consent_table.consent_status IS '許諾ステイタス:request：許諾依頼
consent：許諾済
cancel：許諾解除
denial:否認
expire：許諾期限切れ';
COMMENT ON COLUMN consent_table.consent_flag IS '許諾フラグ';
COMMENT ON COLUMN consent_table.consent_request_date_time IS '許諾依頼日時';
COMMENT ON COLUMN consent_table.consent_end_date_time IS '許諾終了日時';
COMMENT ON COLUMN consent_table.consent_date_time IS '許諾日時';
COMMENT ON COLUMN consent_table.cancel_date_time IS '許諾解除日時';
COMMENT ON COLUMN consent_table.denial_date_time IS '否認日時';

COMMENT ON TABLE rp_service_link_table IS 'RP連携サービス管理テーブル';
COMMENT ON COLUMN rp_service_link_table.service_link_id IS '連携サービスID';
COMMENT ON COLUMN rp_service_link_table.base_id IS 'BaseID';
COMMENT ON COLUMN rp_service_link_table.linking_member_code IS 'サービス連携対象のMember Code';
COMMENT ON COLUMN rp_service_link_table.linking_subsystem_code IS 'サービス連携対象のSubsystem Code';
COMMENT ON COLUMN rp_service_link_table.id_service_code IS 'IDサービスコード';
COMMENT ON COLUMN rp_service_link_table.id_service_userid IS 'IDサービスのユーザID';
COMMENT ON COLUMN rp_service_link_table.uax_link_connect_flag IS 'UAX-LINK連携フラグ';
COMMENT ON COLUMN rp_service_link_table.linking_date_time IS 'サービス連携日時';

COMMENT ON TABLE consent_request_setting_master IS '許諾依頼設定マスタ';
COMMENT ON COLUMN consent_request_setting_master.consent_request_id IS 'id';
COMMENT ON COLUMN consent_request_setting_master.consent_request_setting_id IS '許諾依頼設定id';
COMMENT ON COLUMN consent_request_setting_master.open_flag IS '公開フラグ';
COMMENT ON COLUMN consent_request_setting_master.consent_type IS '許諾依頼タイプ';
COMMENT ON COLUMN consent_request_setting_master.consent_end_date_time IS '許諾終了日時';
COMMENT ON COLUMN consent_request_setting_master.data_provider_member_code IS 'Dataプロバイダのmember code';
COMMENT ON COLUMN consent_request_setting_master.data_provider_subsystem_code IS 'Dataプロバイダのsubsystem code';
COMMENT ON COLUMN consent_request_setting_master.data_provider_service_code IS 'Dataプロバイダのservice code';
COMMENT ON COLUMN consent_request_setting_master.service_provider_member_code IS 'Serviceプロバイダーのmember code';
COMMENT ON COLUMN consent_request_setting_master.service_provider_subsystem_code IS 'Serviceプロバイダーのsubsystem code';
COMMENT ON COLUMN consent_request_setting_master.purpose_of_use IS '利用目的';
COMMENT ON COLUMN consent_request_setting_master.merit_of_consent IS '許諾のメリット';
COMMENT ON COLUMN consent_request_setting_master.create_date IS '作成日';
COMMENT ON COLUMN consent_request_setting_master.update_date IS '更新日';

COMMENT ON TABLE id_convert_master IS 'ID変換マスタ';
COMMENT ON COLUMN id_convert_master.id_convert_id IS 'ID変換ID';
COMMENT ON COLUMN id_convert_master.base_id IS 'BaseID';
COMMENT ON COLUMN id_convert_master.main_id_flag IS 'メインIDフラグ';
COMMENT ON COLUMN id_convert_master.id_service_code IS 'IDサービスコード';
COMMENT ON COLUMN id_convert_master.id_service_userid IS 'IDサービスのユーザID';
COMMENT ON COLUMN id_convert_master.email IS 'IDサービスのユーザメールアドレス';
COMMENT ON COLUMN id_convert_master.email_verified IS 'IDサービスのメールアドレス検証結果';
COMMENT ON COLUMN id_convert_master.ial IS 'IDサービス身元確認保証レベル';
COMMENT ON COLUMN id_convert_master.aal IS 'IDサービス当人認証保証レベル';
COMMENT ON COLUMN id_convert_master.fal IS 'IDサービス認証連携強度レベル';
COMMENT ON COLUMN id_convert_master.linking_date_time IS 'ID連携日時';

COMMENT ON TABLE id_service_master IS 'IDサービスマスタ:';
COMMENT ON COLUMN id_service_master.id_service_id IS 'IDサービスID';
COMMENT ON COLUMN id_service_master.id_service_code IS 'IDサービスコード';
COMMENT ON COLUMN id_service_master.id_service_name IS 'IDサービス名称';
COMMENT ON COLUMN id_service_master.id_service_logo IS 'IDサービスロゴ';
COMMENT ON COLUMN id_service_master.id_provider_name IS 'IDプロバイダー名称';
COMMENT ON COLUMN id_service_master.create_date_time IS '作成日時';
COMMENT ON COLUMN id_service_master.update_date_time IS '更新日時';

COMMENT ON TABLE service_master IS 'データサービスマスタ';
COMMENT ON COLUMN service_master.service_id IS 'サービスid';
COMMENT ON COLUMN service_master.member_code IS 'メンバーコード';
COMMENT ON COLUMN service_master.subsystem_code IS 'サブシステムコード';
COMMENT ON COLUMN service_master.service_code IS 'サービスコード';
COMMENT ON COLUMN service_master.service_version IS 'サービスバージョン';
COMMENT ON COLUMN service_master.service_name IS 'サービス名称';
COMMENT ON COLUMN service_master.description IS '説明';
COMMENT ON COLUMN service_master.input_parameters IS '入力パラメータ';
COMMENT ON COLUMN service_master.output_parameters IS '出力パラメータ';
COMMENT ON COLUMN service_master.create_date IS '作成日';
COMMENT ON COLUMN service_master.update_date IS '更新日';

COMMENT ON TABLE subsystem_master IS 'サブシステムマスタ';
COMMENT ON COLUMN subsystem_master.subsystem_id IS 'サブシステムid';
COMMENT ON COLUMN subsystem_master.member_code IS 'メンバーコード';
COMMENT ON COLUMN subsystem_master.subsystem_code IS 'サブシステムコード';
COMMENT ON COLUMN subsystem_master.subsystem_name IS 'サブシステム名称';
COMMENT ON COLUMN subsystem_master.service_url IS 'サービスURL';
COMMENT ON COLUMN subsystem_master.service_logo_file IS 'サービスロゴ';
COMMENT ON COLUMN subsystem_master.description IS '説明';
COMMENT ON COLUMN subsystem_master.create_date IS '作成日';
COMMENT ON COLUMN subsystem_master.update_date IS '更新日';

COMMENT ON TABLE user_master IS 'ユーザマスタ';
COMMENT ON COLUMN user_master.id IS 'ユーザID';
COMMENT ON COLUMN user_master.base_id IS 'BaseID';
COMMENT ON COLUMN user_master.email IS 'メインemailアドレス';
COMMENT ON COLUMN user_master.main_id_service IS 'メインIDサービス';
COMMENT ON COLUMN user_master.kyc_flag IS '本人確認フラグ';
COMMENT ON COLUMN user_master.latest_ial IS '最新の本人確認保証レベル';
COMMENT ON COLUMN user_master.latest_kyc_date_time IS '最新の本人確認日時';
COMMENT ON COLUMN user_master.mobile_phone IS '携帯電話番号';
COMMENT ON COLUMN user_master.mobile_phone_validation IS '携帯電話番号確認';
COMMENT ON COLUMN user_master.terms_of_service_consent_flag IS '利用規約同意フラグ';
COMMENT ON COLUMN user_master.terms_of_service_consent_date_time IS '利用規約同意日時';
COMMENT ON COLUMN user_master.create_date_time IS '作成日時';
COMMENT ON COLUMN user_master.update_date_time IS '更新日時';

COMMENT ON TABLE member_code_master IS 'メンバーコードマスタ';
COMMENT ON COLUMN member_code_master.member_id IS 'メンバーid';
COMMENT ON COLUMN member_code_master.member_class IS 'メンバークラス';
COMMENT ON COLUMN member_code_master.member_code IS 'メンバーコード';
COMMENT ON COLUMN member_code_master.member_name IS 'メンバー名称';
COMMENT ON COLUMN member_code_master.description IS '説明';
COMMENT ON COLUMN member_code_master.create_date IS '作成日';
COMMENT ON COLUMN member_code_master.update_date IS '更新日';

