from fastapi import APIRouter
from steamapi import SteamAPI
from utils import BannerHelper

client = SteamAPI()

header_router = APIRouter()

banner_helper = BannerHelper()

@header_router.get("/profile/{user_id}")
def get_headers_by_user_id(user_id: str):
    response = client.get_owned_game_ids(user_id)

    response = banner_helper.get_header_urls(response)

    return response

@header_router.get("/profile/{user_id}/completed")
def get_headers_by_user_id_completed(user_id: str):
    response = client.get_completed_game_ids(user_id)

    response = banner_helper.get_header_urls(response)

    return response