from fastapi import APIRouter
from steamapi import SteamAPI
from utils import BannerHelper

client = SteamAPI()

profile_router = APIRouter()

banner_helper = BannerHelper()

@profile_router.get("/{user_id}")
def get_header_by_user_id(bannertype: str, user_id: str, completed: str):
    response = {}

    response['profile'] = client.get_player_profile(user_id)

    game_ids = client.get_owned_game_ids(user_id) if completed == "false" else client.get_completed_game_ids(user_id) if completed == "true" else None

    response['banners'] = banner_helper.get_box_art_urls(game_ids) if bannertype == "boxart" else banner_helper.get_header_urls(game_ids) if bannertype == "header" else banner_helper.get_box_art_urls(game_ids)

    return response