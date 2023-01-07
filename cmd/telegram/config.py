import toml

class Config:
    def __init__(self, config):
        self.Token = config["tg_token"]
        self.Admins = config["tg_admins"]
        self.Addr = "http://" + config["api_host"] + ":" + config["bind_addr"]
        self.JWT = ""

    def update_jwt(self, token):
        self.JWT = token

def get_config(path = "configs/config.toml"):
    string_toml = ""
    with open(path, "r", encoding="UTF-8") as f:
        for row in f.readlines():
            string_toml += row
    
    return Config(toml.loads(string_toml))