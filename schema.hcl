table "users" {
  schema = schema.public
  column "id" {
    null = false
    type = serial
  }
  column "name" {
    null = false
    type = text
  }
  primary_key {
    columns = [column.id]
  }
  index "users_name_idx" {
    on {
      column = "name"
    }
  }
}
schema "public" {
  comment = "Default public gomin schema"
}
