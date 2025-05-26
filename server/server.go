package server

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/mux"
)

type spaHandler struct {
	StaticPath string
	IndexPath  string
}

type BotApi struct {
	session *discordgo.Session
}

func NewBotAPI(s *discordgo.Session) *BotApi {
	return &BotApi{
		session: s,
	}
}

func (api *BotApi) CreateRouter() *mux.Router {
	r := mux.NewRouter()
	// API Endpoints
	r.HandleFunc("/api/health", healthCheck).Methods("GET")
	r.HandleFunc("/api/server/status", serverStatusHandler).Methods("GET")

	// Bot Endpoints
	r.HandleFunc("/api/bot", api.getBotInfo).Methods("GET")
	r.HandleFunc("/api/guilds", api.getGuilds).Methods("GET")
	r.HandleFunc("/api/stats", api.getStats).Methods("GET")

	// Register the SPA handler for serving static files
	spa := &spaHandler{StaticPath: "web/build", IndexPath: "index.html"}
	r.PathPrefix("/").Handler(spa)

	return r
}

func (api *BotApi) getBotInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	botInfo := map[string]interface{}{
		"id":            api.session.State.User.ID,
		"username":      api.session.State.User.Username,
		"discriminator": api.session.State.User.Discriminator,
		"avatar":        api.session.State.User.Avatar,
		"bot":           api.session.State.User.Bot,
	}

	json.NewEncoder(w).Encode(botInfo)
}

func (api *BotApi) getGuilds(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	guilds := make([]map[string]interface{}, 0)

	for _, guild := range api.session.State.Guilds {
		guildInfo := map[string]interface{}{
			"id":           guild.ID,
			"name":         guild.Name,
			"icon":         guild.Icon,
			"member_count": guild.MemberCount,
			"owner_id":     guild.OwnerID,
			"joined_at":    guild.JoinedAt,
		}
		guilds = append(guilds, guildInfo)
	}

	json.NewEncoder(w).Encode(guilds)
}

func (api *BotApi) getStats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	totalMembers := 0
	for _, guild := range api.session.State.Guilds {
		totalMembers += guild.MemberCount
	}

	// Replace this later with actual uptime tracking
	// For now, we will just use the current time as a placeholder
	timeSince := time.Now()

	stats := map[string]interface{}{
		"guild_count":   len(api.session.State.Guilds),
		"total_members": totalMembers,
		"uptime":        time.Since(timeSince).String(),
	}

	json.NewEncoder(w).Encode(stats)
}

func serverStatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func (h *spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(h.StaticPath, r.URL.Path)

	fi, err := os.Stat(path)
	if os.IsNotExist(err) || fi.IsDir() {
		http.ServeFile(w, r, filepath.Join(h.StaticPath, h.IndexPath))
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.FileServer(http.Dir(h.StaticPath)).ServeHTTP(w, r)
}
