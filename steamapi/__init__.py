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
            res = requests.get(url).json()

            game_ids = [game['appid'] for game in res['response']['games']]

            return game_ids
        except OSError:
            return OSError
    
    def get_completed_game_data(self, steam_id: str):
        try:
            game_ids = self.get_owned_game_ids(steam_id=steam_id)

            completed_game_ids = []

            # Not practical. The Steam Web API has forced my hand.
            for id in game_ids:
                url = f'http://api.steampowered.com/ISteamUserStats/GetPlayerAchievements/v0001/?appid={id}&key={self.api_key}&steamid={steam_id}'
                res = requests.get(url).json()

                if 'achievements' in res['playerstats']:
                    self.validate_completion_status(res['playerstats']['achievements']) and completed_game_ids.append(id)
            
            return completed_game_ids
        except OSError:
            return OSError
    
    def validate_completion_status(self, achievements: list):
        return all(achievement['achieved'] == 1 for achievement in achievements)
