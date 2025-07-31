from fastapi import APIRouter
from steamapi import SteamAPI
from utils import BannerHelper

client = SteamAPI()

profile_router = APIRouter()

banner_helper = BannerHelper()

@profile_router.get("/{banner_type}/{user_id}")
def get_header_by_user_id(banner_type: str, user_id: str, completed: str):

    response = client.get_owned_game_ids(user_id) if completed == "false" else client.get_completed_game_ids(user_id) if completed == "true" else None

    response = banner_helper.get_box_art_urls(response) if banner_type == "box-art" else banner_helper.get_header_urls(response) if banner_type == "header" else banner_helper.get_box_art_urls(response)

    return response