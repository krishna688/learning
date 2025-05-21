output "file_id" {
  value       = local_file.my_first_file.id
  description = "my_first_file_id"
}

output "file_content" {
  value       = local_file.my_first_file.content
  description = "my_first_file_id content"
}

output "file_256sha" {
  value       = local_file.my_first_file.content_sha256
  description = "my_first_file_id sha"

}