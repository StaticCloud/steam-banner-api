from dotenv import load_dotenv
import requests
import os

load_dotenv()

class SteamAPI():
    def __init__(self):
        self.api_key = os.getenv("API_KEY")

    def get_owned_game_ids(self, steam_id: str):
        try:
            url = f'http://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key={self.api_key}&steamid={steam_id}&format=json'
            res = requests.get(url)

            res = res.json()

            game_ids = [game['appid'] for game in res['response']['games']]
            
            return game_ids
        except OSError:
            return OSError
    
    def get_achievement_data(self, steam_id: str):
        pass
