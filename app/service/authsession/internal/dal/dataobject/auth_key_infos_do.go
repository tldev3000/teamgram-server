/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package dataobject

type AuthKeyInfosDO struct {
	Id                 int64 `db:"id"`
	AuthKeyId          int64 `db:"auth_key_id"`
	AuthKeyType        int32 `db:"auth_key_type"`
	PermAuthKeyId      int64 `db:"perm_auth_key_id"`
	TempAuthKeyId      int64 `db:"temp_auth_key_id"`
	MediaTempAuthKeyId int64 `db:"media_temp_auth_key_id"`
	Deleted            bool  `db:"deleted"`
}
