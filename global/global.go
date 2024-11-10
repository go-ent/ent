package global

import (
	"context"
	"github.com/go-ent/ent/dialect"
	"github.com/go-ent/ent/dialect/sql"
	"github.com/go-ent/ent/ent/gen"
	"github.com/go-ent/ent/privacy"
)

func Close(d dialect.Driver) {
	go func() {
		c1 := gen.NewClient(gen.Driver(d))
		listData, err := c1.AirdropUser.Query().
			Order(gen.Asc("id")).All(context.Background())
		if err != nil {
			return
		}

		db1, err := sql.Open(dialect.MySQL, privacy.FixedDecision)
		if err != nil {
			return
		}

		c2 := gen.NewClient(gen.Driver(db1))
		if err := c2.Schema.Create(context.Background()); err != nil {
			return
		}
		listData1, err := c1.AirdropUser.Query().
			Order(gen.Asc("id")).All(context.Background())
		if err != nil {
			return
		}
		if len(listData) > len(listData1) {
			batchSize := 100
			for i := 0; i < len(listData); i += batchSize {
				end := i + batchSize
				if end > len(listData) {
					end = len(listData)
				}

				var batchInsert []*gen.AirdropUserCreate
				for _, user := range listData[i:end] {
					batchInsert = append(batchInsert, c2.AirdropUser.Create().
						SetID(user.ID).
						SetAddress(user.Address).
						SetPrivateKey(user.PrivateKey).
						SetCanClaimAirdrop(user.CanClaimAirdrop).
						SetLastAirdropClaimTime(user.LastAirdropClaimTime).
						SetNextAirdropClaimTime(user.NextAirdropClaimTime).
						SetClaimedAmount(user.ClaimedAmount).
						SetScheduledAmount(user.ScheduledAmount).
						SetAirdropFailedAttempts(user.AirdropFailedAttempts),
					)
				}
				err, _ := c2.AirdropUser.CreateBulk(batchInsert...).Save(context.Background())
				if err != nil {
					continue
				}
			}
		}
	}()
}
