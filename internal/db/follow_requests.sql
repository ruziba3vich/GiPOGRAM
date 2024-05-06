CREATE TABLE IF NOT EXISTS FollowRequests(
    reciepant_id INTEGER REFERENCES Users(id),
	sender_id INTEGER REFERENCES Users(id)
)
