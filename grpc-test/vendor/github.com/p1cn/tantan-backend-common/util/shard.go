package util

import (
	"strconv"

	slog "github.com/p1cn/tantan-backend-common/log"
)

func PartitionByMod(uids []string, mod int) (map[int][]string, error) {
	res := make(map[int][]string)
	for _, u := range uids {
		id, err := strconv.Atoi(u)
		if err != nil {
			return nil, err
		}
		key := Int32Abs(id) % mod
		if _, find := res[key]; find {
			res[key] = append(res[key], u)
		} else {
			res[key] = []string{u}
		}
	}
	return res, nil
}

func IdToAbsInt(requestID string) (int, error) {
	id, err := strconv.Atoi(requestID)
	if err != nil {
		return 0, err
	}
	id = Int32Abs(id)
	return id, nil
}

func parseShardIdFromMessageID(oId string, mod int) (int64, int, int, error) {
	id, err := strconv.ParseInt(oId, 10, 64)
	if err != nil {
		slog.Err("%v", err)
		return -1, -1, -1, err
	}
	timestamp := (id >> 23) + 1314220021721

	shardId := (id << 41) >> 51

	if shardId <= 0 {
		shardId = int64(mod) + shardId
	}

	seqId := (id << 54) >> 54
	if seqId <= 0 {
		seqId = 1024 + seqId
	}

	return timestamp, int(shardId - 1), int(seqId - 1), nil
}



func MapMessageIdsByShardId(oIds []string, mod int) (map[int][]string, error) {
	idsByShards := make(map[int][]string)
	for _, oId := range oIds {
		_, shardId, _, err := parseShardIdFromMessageID(oId, mod)
		if err != nil {
			return nil, err
		}
		idsByShards[shardId] = append(idsByShards[shardId], oId)
	}
	return idsByShards, nil
}

func MapUserIdsByShardId(userIds []string, mod int) (map[int][]string, error) {
	idsByShards := make(map[int][]string)
	for _, userId := range userIds {
		id, err := strconv.Atoi(userId)
		if err != nil {
			slog.Err("%v", err)
			return nil, err
		}
		shardId := Int32Abs(id) % mod
		idsByShards[shardId] = append(idsByShards[shardId], userId)
	}
	return idsByShards, nil
}

func MapMomentIdsByShardId(oIds []string, mod int) (map[int][]string, error) {
	return MapMessageIdsByShardId(oIds, mod)
}

func ParseShardIdFromMomentId(oId string, mod int) (int64, int, int, error) {
	return  parseShardIdFromMessageID(oId, mod)
}
