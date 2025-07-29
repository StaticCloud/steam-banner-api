from fastapi import FastAPI
from steamapi import SteamAPI

app = FastAPI()

client = SteamAPI()

@app.get("/{user_id}")
def get_user_achievements(user_id: str):
    response = client.get_owned_game_ids(user_id)
    return response

