import random
import string
from datetime import datetime

import lorem
import names
import os
import psycopg2
import sys


def get_db():
    host = os.environ["PG_HOST"]
    port = os.environ["PG_PORT"]
    database = os.environ["PG_DATABASE"]
    user = os.environ["PG_USER"]
    password = os.environ["PG_PASSWORD"]
    conn_str = f"host={host} port={port} dbname={database} user={user} password={password}"
    return psycopg2.connect(conn_str)


def generate_users(conn):
    users = "INSERT INTO users (username, password) VALUES "
    letters = string.ascii_lowercase

    for i in range(3):
        username = names.get_first_name()
        password = ''.join(random.choice(letters) for _ in range(32))
        users += f"('{username}', '{password}')"
        users += ", " if i < 2 else ""
    users += "RETURNING id"
    cur = conn.cursor()
    cur.execute(users)
    conn.commit()
    res = cur.fetchall()
    cur.close()
    return [r[0] for r in res]


def generate_messages(conn, channel: int, senders: list, no_of_messages: int):
    messages = "INSERT INTO messages (sender, time_sent, channel, content) VALUES "
    for i in range(no_of_messages):
        message = lorem.sentence()
        user = random.choice(senders)
        time_sent = str(datetime.utcnow())
        messages += f"('{user}', '{time_sent}', '{channel}', '{message}')"
        messages += ", " if i < no_of_messages - 1 else ""
    cur = conn.cursor()
    cur.execute(messages)
    conn.commit()
    cur.close()


if __name__ == "__main__":
    "Example of call: python3 generate_conversation.py channel=5 no_of_messages=100"
    bash_args = {s[0]: s[1] for s in [a.split('=') for a in sys.argv if '=' in a]}
    db = get_db()
    user_ids = generate_users(conn=db)
    generate_messages(
        conn=db,
        channel=int(bash_args.get("channel", 9)),
        senders=user_ids,
        no_of_messages=int(bash_args.get("no_of_messages", 10))
    )
    db.close()
