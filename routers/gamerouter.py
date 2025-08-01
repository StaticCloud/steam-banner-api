from fastapi import APIRouter
from steamapi import SteamAPI
from utils import BannerHelper
from pydantic import BaseModel

client = SteamAPI()

game_router = APIRouter()

banner_helper = BannerHelper()

class Body(BaseModel):
    game_ids: list[int]

@game_router.get("/")
def get_header_by_game_id(body: Body, bannertype: str):
    return banner_helper.get_box_art_urls(body.game_ids) if bannertype == "boxart" else banner_helper.get_header_urls(body.game_ids) if bannertype == "header" else banner_helper.get_box_art_urls(body.game_ids)