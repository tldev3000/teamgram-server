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

type ChatParticipantsDO struct {
	Id              int64  `db:"id"`
	ChatId          int64  `db:"chat_id"`
	UserId          int64  `db:"user_id"`
	ParticipantType int32  `db:"participant_type"`
	Link            string `db:"link"`
	Usage2          int32  `db:"usage2"`
	AdminRights     int32  `db:"admin_rights"`
	InviterUserId   int64  `db:"inviter_user_id"`
	InvitedAt       int64  `db:"invited_at"`
	KickedAt        int64  `db:"kicked_at"`
	LeftAt          int64  `db:"left_at"`
	State           int32  `db:"state"`
	Date2           int64  `db:"date2"`
}
