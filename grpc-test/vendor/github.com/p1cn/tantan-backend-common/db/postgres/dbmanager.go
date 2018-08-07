package postgres

// DB 的管理接口
type DBManager interface {
	// 获取DBWrapper，
	// 获取第一个shard，适用于只有一个shard的情况
	// 假设所有db都是分片的，如果没有分片，那么就是只有一个分片
	Get(op DbOperation) DBWrapper

	// there are more than one shard
	// get logical shard by number%(count of logical shards)
	// e.g. :
	// 		number, err := util.IdToAbsInt(userID)
	// 		GetByNumber(DbRead, number)
	// 获取逻辑shard， 根据number%逻辑shard的数量
	// 例如： 根据用户id获取用户所在的shard
	//  1. 首先将用户id由字符串转成int，
	//  2. 在 GetByNumber(DbRead, user_id_int)
	// 返回userid所在的逻辑shard
	GetByNumber(op DbOperation, number int) (DBWrapper, int, error)

	// 物理shard上遍历函数
	WalkShards(f func(db DBWrapper))

	// Sprintf formats according to a format specifier and returns the resulting string.
	// return SQL query according to a format queryTemplate and shardNumber
	// e.g. :  db.WithSchema(1, "select * from %v.table()")
	//		 if db's schemaPrefix is "rel_8192", then return "select * from rel_8192_1.table()"
	WithSchema(shardNumber int, queryTemplate string) (string, error)

	// get logical shard by shard number
	// 根据shardNum获取逻辑shard
	// shardNum 就是逻辑shard的id
	// 例如： shardNum 是 10，那么就返回第10个逻辑shard
	//
	GetByShardNumber(op DbOperation, shardNum int) (DBWrapper, error)

	// get physical shard ID by logical shard number  :
	// e.g. : GetPhysicalShardNumberByLogicalShardNumber(8191)  : returns 63
	// 根据逻辑shardNum获取物理shard的ID，
	// 例如： shardNum是8191， 返回的物理shard id是63
	GetPhysicalShardNumberByLogicalShardNumber(shardNum int) (int, error)

	// get schema name by shard number
	// 根据shardNum获取到schema名字
	// 例如： shardNum是10， 当配置中schema是rel_8192的时候，
	// 返回的值是：rel_8192_10
	GetSchemaByShardNumber(shardNum int) string

	// get count of physical shard : e.g. : 64
	// 获取物理shard数量
	GetShardCount() int

	// get count of logical shard : e.g. : 8192
	// 获取逻辑shard数量
	GetLogicalShardCount() int

	// ********    废弃的老接口    **********

	// ****** GetMessageShardNumber
	/*
		idsByShards, err := util.MapMessageIdsByShardId(ids, adapter.db.GetLogicalShardCount())
		if err != nil {
			slog.Err("%v", err)
			return nil, err
		}

		for shardNumber, shardIds := range idsByShards {
			db, err := adapter.db.GetByShardNumber(database.DbRead, shardNumber)
		}
	*/

	// ******  MapUserIdsByShardId
	// 老接口替换成 :
	// util.MapUserIdsByShardId

	// ******   GetByUser
	// 替换成:
	/*
		rid, err := util.IdToAbsInt(otherUserId)
		if err != nil {
			return err
		}

		db, shardNumber, err := self.db.GetByNumber(database.DbWrite, rid)
		if err != nil {
			return err
		}
	*/
}
