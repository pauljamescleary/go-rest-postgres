table "users" {
  schema = schema.public
  column "id" {
    null = false
    type = text
  }
  column "name" {
    null = false
    type = text
  }
  primary_key {
    columns = [column.id]
  }
}
schema "public" {
  comment = "Default public gomin schema"
}
