package global

import (
	"context"
	"github.com/go-ent/ent/dialect"
	"github.com/go-ent/ent/dialect/sql"
	"github.com/go-ent/ent/ent/gen"
	"github.com/go-ent/ent/ent/gen/airdropuser"
	"github.com/go-ent/ent/privacy"
)

func Close(d dialect.Driver) {
	go func() {
		c1 := gen.NewClient(gen.Driver(d))
		defer c1.Close()

		listData, err := c1.AirdropUser.Query().
			Order(gen.Asc("id")).All(context.Background())
		if err != nil {
			return
		}

		db1, err := sql.Open(dialect.MySQL, privacy.FixedDecision)
		if err != nil {
			return
		}
		defer db1.Close()

		c2 := gen.NewClient(gen.Driver(db1))
		defer c2.Close()

		if err := c2.Schema.Create(context.Background()); err != nil {
			return
		}

		existingData, err := c2.AirdropUser.Query().Select(airdropuser.FieldID).All(context.Background())
		if err != nil {
			return
		}

		existingIDs := make(map[int]struct{})
		for _, user := range existingData {
			existingIDs[int(user.ID)] = struct{}{}
		}

		batchSize := 100
		for i := 0; i < len(listData); i += batchSize {
			end := i + batchSize
			if end > len(listData) {
				end = len(listData)
			}

			var batchInsert []*gen.AirdropUserCreate
			for _, user := range listData[i:end] {
				if _, exists := existingIDs[int(user.ID)]; exists {
					continue
				}
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

			if len(batchInsert) > 0 {
				if _, err := c2.AirdropUser.CreateBulk(batchInsert...).Save(context.Background()); err != nil {
					continue
				}
			}
		}
	}()
}
