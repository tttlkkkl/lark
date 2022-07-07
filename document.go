package lark

type CopyRequest struct {
	CommentNeeded  bool    `json:"commentNeeded"`
	DstFolderToken string  `json:"dstFolderToken"`
	DstName        string  `json:"dstName"`
	Type           DocType `json:"type"`
}

type CopyResponse struct {
	Response
	Data struct {
		FolderToken string  `json:"folderToken"`
		Revision    int64   `json:"revision"`
		Token       string  `json:"token"`
		Type        DocType `json:"type"`
		URL         string  `json:"url"`
	} `json:"data"`
}

type UpdatePermissionRequest struct {
	CommentEntity   string `json:"comment_entity"`
	ExternalAccess  bool   `json:"external_access"`
	InviteExternal  bool   `json:"invite_external"`
	LinkShareEntity string `json:"link_share_entity"`
	SecurityEntity  string `json:"security_entity"`
	ShareEntity     string `json:"share_entity"`
}

type UpdatePermissionResponse struct {
	Response
	Data struct {
		PermissionPublic struct {
			CommentEntity   string `json:"comment_entity"`
			ExternalAccess  bool   `json:"external_access"`
			InviteExternal  bool   `json:"invite_external"`
			LinkShareEntity string `json:"link_share_entity"`
			SecurityEntity  string `json:"security_entity"`
			ShareEntity     string `json:"share_entity"`
		} `json:"permission_public"`
	} `json:"data"`
}
