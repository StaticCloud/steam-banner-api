from fastapi import FastAPI
from steamapi import SteamAPI

app = FastAPI()

client = SteamAPI()

@app.get("/games/{user_id}")
def get_owned_game_ids(user_id: str):
    response = client.get_owned_game_ids(user_id)

    return response

@app.get("/games/{user_id}/completed")
def get_completed_game_ids(user_id: str):
    response = client.get_completed_game_ids(user_id)

    return response