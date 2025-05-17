table "refresh_tokens" {
  schema = schema.public
  column "user_id" {
    null = false
    type = uuid
  }
  column "hashed_token" {
    null = false
    type = text
  }
  column "created_at" {
    null    = false
    type    = timestamptz
    default = sql("CURRENT_TIMESTAMP")
  }
  column "expires_at" {
    null    = false
    type    = timestamptz
    default = sql("(CURRENT_TIMESTAMP + '1 day'::interval)")
  }
  primary_key {
    columns = [column.user_id, column.hashed_token]
  }
  foreign_key "refresh_tokens_user_id_fkey" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  unique "refresh_tokens_hashed_token_key" {
    columns = [column.hashed_token]
  }
}
table "reports" {
  schema = schema.public
  column "user_id" {
    null = false
    type = uuid
  }
  column "id" {
    null    = false
    type    = uuid
    default = sql("gen_random_uuid()")
  }
  column "report_type" {
    null = false
    type = text
  }
  column "output_file_path" {
    null = true
    type = text
  }
  column "download_url" {
    null = true
    type = text
  }
  column "download_url_expires_at" {
    null = true
    type = timestamptz
  }
  column "error_message" {
    null = true
    type = text
  }
  column "created_at" {
    null    = false
    type    = timestamptz
    default = sql("CURRENT_TIMESTAMP")
  }
  column "started_at" {
    null = true
    type = timestamptz
  }
  column "failed_at" {
    null = true
    type = timestamptz
  }
  primary_key {
    columns = [column.user_id, column.id]
  }
  foreign_key "reports_user_id_fkey" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
}
table "users" {
  schema = schema.public
  column "id" {
    null    = false
    type    = uuid
    default = sql("gen_random_uuid()")
  }
  column "email" {
    null = false
    type = text
  }
  column "hashed_password" {
    null = false
    type = text
  }
  column "created_at" {
    null    = false
    type    = timestamptz
    default = sql("CURRENT_TIMESTAMP")
  }
  primary_key {
    columns = [column.id]
  }
  unique "users_email_key" {
    columns = [column.email]
  }
}
schema "public" {
  comment = "standard public schema"
}
