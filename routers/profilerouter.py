from fastapi import APIRouter, HTTPException, Query
from steamapi import SteamAPI
from utils import BannerHelper

client = SteamAPI()

profile_router = APIRouter()

banner_helper = BannerHelper()

@profile_router.get("/{user_id}")
def get_header_by_user_id(
        bannertype: str = Query(..., regex="^(boxart|header)$"), 
        user_id: str = None, 
        completed: str = Query(..., regex="^(true|false)")
    ):

    if not user_id:
        raise HTTPException(status_code=400, detail="A Steam user id was not provided.")

    try: 
        profile = client.get_player_profile(user_id)
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Error fetching profile: {str(e)}")

    if profile is None:
        raise HTTPException(status_code=404, detail="Could not locate a Steam user with the provided id.")

    try:
        if completed == "true":
            game_ids = client.get_completed_game_ids(user_id)
        else:
            game_ids = client.get_owned_game_ids(user_id)
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Error fetching banners: {str(e)}")
    
    try:
        if bannertype == "header":
            banners = banner_helper.get_header_urls(game_ids)
        else:
            banners = banner_helper.get_box_art_urls(game_ids)
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Error fetching banners: {str(e)}")

    return {
        "profile": profile,
        "banners": banners
    }