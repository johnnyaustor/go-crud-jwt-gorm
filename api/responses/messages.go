package responses

const (
	MessageAuthorizedUserPutData = 103
	MessageAuthorizedUserDeleteData = 104
	MessageAuthorizedPostsCreateData = 111
	MessageAuthorizedPostsUpdateData = 113
	MessageAuthorizedPostsDeleteData = 114
	MessagePostNotFound = 110
)

var messages = map[int]string {
	MessageAuthorizedUserPutData: "Kamu tidak diperbolehkan mengubah data user lain",
	MessageAuthorizedUserDeleteData: "Kamu tidak diperbolehkan menghapus data user lain",
	MessageAuthorizedPostsCreateData: "Kamu tidak diperbolehkan menambahkan post data user lain",
	MessageAuthorizedPostsUpdateData: "Kamu tidak diperbolehkan mengubah post data user lain",
	MessageAuthorizedPostsDeleteData: "Kamu tidak diperbolehkan menghapus post data user lain",
	MessagePostNotFound: "Post not found",
}

func GetMessage(code int) string {
	return messages[code]
}