from fastapi import APIRouter
from steamapi import SteamAPI

client = SteamAPI()

router = APIRouter()

@router.get("/{user_id}")
def get_owned_game_ids(user_id: str):
    response = client.get_owned_game_ids(user_id)

    return response

@router.get("/{user_id}/completed")
def get_completed_game_ids(user_id: str):
    response = client.get_completed_game_ids(user_id)

    return response