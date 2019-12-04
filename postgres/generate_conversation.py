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


def generate_messages(conn, channel: str, senders: list, no_of_messages: int):
    messages = "INSERT INTO messages (sender, date, room, text) VALUES "
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
    generate_messages(
        conn=db,
        channel=bash_args.get("channel", 9),
        senders=["xxx", "yyy", "zzz"],
        no_of_messages=int(bash_args.get("no_of_messages", 10))
    )
    db.close()
