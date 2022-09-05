package puzzad

import (
	"testing"
)

func TestHashAndVerify(t *testing.T) {
	tests := []struct {
		name    string
		plain1  string
		plain2  string
		want    bool
		wantErr bool
	}{
		{
			name:    "VerifyCorrect",
			plain1:  "test",
			plain2:  "test",
			want:    true,
			wantErr: false,
		},
		{
			name:    "VerifyIncorrect",
			plain1:  "test",
			plain2:  "test2",
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetHash(tt.plain1)
			if err != nil {
				t.Errorf("CheckHash() error = %v", err)
				return
			}
			got2, err := CheckHash(tt.plain2, got)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != got2 {
				t.Errorf("CheckHash() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckHash(t *testing.T) {
	tests := []struct {
		name    string
		plain   string
		hash    string
		want    bool
		wantErr bool
	}{
		{
			name:    "Correct password for hash",
			plain:   "test",
			hash:    "argon2id$65536$1$4$aJuLQpIeRmFotQzV90bnIMe7e49RMTm5Ve8g89xhUrQ$DiASQQu2g3NneKskie+z9B2lUFSvD7s5ZN/ggPGR8arPPXpV+ZJPr1HZm15Z8LPgGsIpqBcNG8hj3nACLkkZdA",
			want:    true,
			wantErr: false,
		},
		{
			name:    "Incorrect password for hash",
			plain:   "test2",
			hash:    "argon2id$65536$1$4$aJuLQpIeRmFotQzV90bnIMe7e49RMTm5Ve8g89xhUrQ$DiASQQu2g3NneKskie+z9B2lUFSvD7s5ZN/ggPGR8arPPXpV+ZJPr1HZm15Z8LPgGsIpqBcNG8hj3nACLkkZdA",
			want:    false,
			wantErr: false,
		},
		{
			name:    "Invalid hash type",
			plain:   "test2",
			hash:    "argon2$65536$1$4$aJuLQpIeRmFotQzV90bnIMe7e49RMTm5Ve8g89xhUrQ$DiASQQu2g3NneKskie+z9B2lUFSvD7s5ZN/ggPGR8arPPXpV+ZJPr1HZm15Z8LPgGsIpqBcNG8hj3nACLkkZdA",
			want:    false,
			wantErr: true,
		},
		{
			name:    "Invalid hash",
			plain:   "test2",
			hash:    "argon2id$65536$1$4$aJuLQpIeRmFotQzV90bnIMe7e49RMTm5Ve8g89xhUrQ",
			want:    false,
			wantErr: true,
		},
		{
			name:    "Bad Memory",
			plain:   "test2",
			hash:    "argon2id$$1$4$aJuLQpIeRmFotQzV90bnIMe7e49RMTm5Ve8g89xhUrQ$DiASQQu2g3NneKskie+z9B2lUFSvD7s5ZN/ggPGR8arPPXpV+ZJPr1HZm15Z8LPgGsIpqBcNG8hj3nACLkkZdA",
			want:    false,
			wantErr: true,
		},
		{
			name:    "Bad Time",
			plain:   "test2",
			hash:    "argon2id$1$a$4$aJuLQpIeRmFotQzV90bnIMe7e49RMTm5Ve8g89xhUrQ$DiASQQu2g3NneKskie+z9B2lUFSvD7s5ZN/ggPGR8arPPXpV+ZJPr1HZm15Z8LPgGsIpqBcNG8hj3nACLkkZdA",
			want:    false,
			wantErr: true,
		},
		{
			name:    "Bad Threads",
			plain:   "test2",
			hash:    "argon2id$1$1$a$aJuLQpIeRmFotQzV90bnIMe7e49RMTm5Ve8g89xhUrQ$DiASQQu2g3NneKskie+z9B2lUFSvD7s5ZN/ggPGR8arPPXpV+ZJPr1HZm15Z8LPgGsIpqBcNG8hj3nACLkkZdA",
			want:    false,
			wantErr: true,
		},
		{
			name:    "Bad Salt",
			plain:   "test2",
			hash:    "argon2id$1$1$a$aJuLQpIeRmFotQzV90bnIMe7e49RMT^m5Ve8g89xhUrQ$DiASQQu2g3NneKskie+z9B2lUFSvD7s5ZN/ggPGR8arPPXpV+ZJPr1HZm15Z8LPgGsIpqBcNG8hj3nACLkkZdA",
			want:    false,
			wantErr: true,
		},
		{
			name:    "empty Salt",
			plain:   "test2",
			hash:    "argon2id$1$1$a$$DiASQQu2g3NneKskie+z9B2lUFSvD7s5ZN/ggPGR8arPPXpV+ZJPr1HZm15Z8LPgGsIpqBcNG8hj3nACLkkZdA",
			want:    false,
			wantErr: true,
		},
		{
			name:    "Bad Hash",
			plain:   "test2",
			hash:    "argon2id$1$1$a$aJuLQpIeRmFotQzV90bnIMe7e49RMTm5Ve8g89xhUrQ$DiASQQu2g3NneKskie+z9B2lUFSvD7s5ZN/ggPGR8arPPXpV+ZJPr1HZm15Z8LPgGsIpqBcNG8hj3nACLkkZdA",
			want:    false,
			wantErr: true,
		},
		{
			name:    "Empty Hash",
			plain:   "test2",
			hash:    "argon2id$1$1$a$aJuLQpIeRmFotQzV90bnIMe7e49RMTm5Ve8g89xhUrQ$",
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckHash(tt.plain, tt.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckHash() got = %v, want %v", got, tt.want)
			}
		})
	}
}
