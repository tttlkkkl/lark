package lark

import (
	"fmt"
	"testing"
)

func TestAESDecrypt(t *testing.T) {
	type args struct {
		base64CipherText string
		key              []byte
	}
	tests := []struct {
		name               string
		args               args
		wantUnpadDecrypted []byte
		wantErr            bool
	}{
		{
			name: "",
			args: args{
				base64CipherText: "FIAfJPGRmFZWkaxPQ1XrJZVbv2JwdjfLk4jx0k/U1deAqYK3AXOZ5zcHt/cC4ZNTqYwWUW/EoL+b2hW/C4zoAQQ5CeMtbxX2zHjm+E4nX/Aww+FHUL6iuIMaeL2KLxqdtbHRC50vgC2YI7xohnb3KuCNBMUzLiPeNIpVdnYaeteCmSaESb+AZpJB9PExzTpRDzCRv+T6o5vlzaE8UgIneC1sYu85BnPBEMTSuj1ZZzfdQi7ZW992Z4dmJxn9e8FL2VArNm99f5Io3c2O4AcNsQENNKtfAAxVjCqc3mg5jF0nKabA+u/5vrUD76flX1UOF5fzJ0sApG2OEn9wfyPDRBsApn9o+fceF9hNrYBGsdtZrZYyGG387CGOtKsuj8e2E8SNp+Pn4E9oYejOTR+ZNLNi+twxaXVlJhr6l+RXYwEiMGQE9zGFBD6h2dOhKh3W84p1GEYnSRIz1+9/Hp66arjC7RCrhuW5OjCj4QFEQJiwgL45XryxHtiZ7JdAlPmjVsL03CxxFZarzxzffryrWUG3VkRdHRHbTsC34+ScoL5MTDU1QAWdqUC1T7xT0lCvQELaIhBTXAYrznJl6PlA83oqlMxpHh0gZBB1jFbfoUr7OQbBs1xqzpYK6Yjux6diwpQB1zlZErYJUfCqK7G/zI9yK/60b4HW0k3M+AvzMcw=",
				key:              []byte("kudryavka"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUnpadDecrypted, err := AESDecrypt(tt.args.base64CipherText, tt.args.key)
			fmt.Println("---->", string(gotUnpadDecrypted), err)
			// if (err != nil) != tt.wantErr {
			// 	t.Errorf("AESDecrypt() error = %v, wantErr %v", err, tt.wantErr)
			// 	return
			// }
			// if !reflect.DeepEqual(gotUnpadDecrypted, tt.wantUnpadDecrypted) {
			// 	t.Errorf("AESDecrypt() = %v, want %v", gotUnpadDecrypted, tt.wantUnpadDecrypted)
			// }
		})
	}
}
