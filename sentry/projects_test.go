package sentry

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectService_List(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/0/projects/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[
			{
				"status": "active",
				"slug": "the-spoiled-yoghurt",
				"defaultEnvironment": null,
				"features": [
					"data-forwarding",
					"rate-limits"
				],
				"color": "#bf6e3f",
				"isPublic": false,
				"dateCreated": "2017-07-18T19:29:44.996Z",
				"platforms": [],
				"callSign": "THE-SPOILED-YOGHURT",
				"firstEvent": null,
				"processingIssues": 0,
				"organization": {
					"name": "The Interstellar Jurisdiction",
					"slug": "the-interstellar-jurisdiction",
					"avatar": {
						"avatarUuid": null,
						"avatarType": "letter_avatar"
					},
					"dateCreated": "2017-07-18T19:29:24.565Z",
					"id": "2",
					"isEarlyAdopter": false
				},
				"isBookmarked": false,
				"callSignReviewed": false,
				"id": "4",
				"name": "The Spoiled Yoghurt"
			},
			{
				"status": "active",
				"slug": "prime-mover",
				"defaultEnvironment": null,
				"features": [
					"data-forwarding",
					"rate-limits",
					"releases"
				],
				"color": "#bf5b3f",
				"isPublic": false,
				"dateCreated": "2017-07-18T19:29:30.063Z",
				"platforms": [],
				"callSign": "PRIME-MOVER",
				"firstEvent": null,
				"processingIssues": 0,
				"organization": {
					"name": "The Interstellar Jurisdiction",
					"slug": "the-interstellar-jurisdiction",
					"avatar": {
						"avatarUuid": null,
						"avatarType": "letter_avatar"
					},
					"dateCreated": "2017-07-18T19:29:24.565Z",
					"id": "2",
					"isEarlyAdopter": false
				},
				"isBookmarked": false,
				"callSignReviewed": false,
				"id": "3",
				"name": "Prime Mover"
			},
			{
				"status": "active",
				"slug": "pump-station",
				"defaultEnvironment": null,
				"features": [
					"data-forwarding",
					"rate-limits",
					"releases"
				],
				"color": "#3fbf7f",
				"isPublic": false,
				"dateCreated": "2017-07-18T19:29:24.793Z",
				"platforms": [],
				"callSign": "PUMP-STATION",
				"firstEvent": null,
				"processingIssues": 0,
				"organization": {
					"name": "The Interstellar Jurisdiction",
					"slug": "the-interstellar-jurisdiction",
					"avatar": {
						"avatarUuid": null,
						"avatarType": "letter_avatar"
					},
					"dateCreated": "2017-07-18T19:29:24.565Z",
					"id": "2",
					"isEarlyAdopter": false
				},
				"isBookmarked": false,
				"callSignReviewed": false,
				"id": "2",
				"name": "Pump Station"
			}
		]`)
	})

	client := NewClient(httpClient, nil, "")
	projects, _, err := client.Projects.List()
	assert.NoError(t, err)

	expectedOrganization := Organization{
		ID:          "2",
		Slug:        "the-interstellar-jurisdiction",
		Name:        "The Interstellar Jurisdiction",
		DateCreated: mustParseTime("2017-07-18T19:29:24.565Z"),
		Avatar: OrganizationAvatar{
			Type: "letter_avatar",
		},
		IsEarlyAdopter: false,
	}
	expected := []Project{
		{
			ID:           "4",
			Slug:         "the-spoiled-yoghurt",
			Name:         "The Spoiled Yoghurt",
			DateCreated:  mustParseTime("2017-07-18T19:29:44.996Z"),
			IsPublic:     false,
			IsBookmarked: false,
			CallSign:     "THE-SPOILED-YOGHURT",
			Color:        "#bf6e3f",
			Features: []string{
				"data-forwarding",
				"rate-limits",
			},
			Status:       "active",
			Organization: expectedOrganization,
		},
		{
			ID:           "3",
			Slug:         "prime-mover",
			Name:         "Prime Mover",
			DateCreated:  mustParseTime("2017-07-18T19:29:30.063Z"),
			IsPublic:     false,
			IsBookmarked: false,
			CallSign:     "PRIME-MOVER",
			Color:        "#bf5b3f",
			Features: []string{
				"data-forwarding",
				"rate-limits",
				"releases",
			},
			Status:       "active",
			Organization: expectedOrganization,
		},
		{
			ID:           "2",
			Slug:         "pump-station",
			Name:         "Pump Station",
			DateCreated:  mustParseTime("2017-07-18T19:29:24.793Z"),
			IsPublic:     false,
			IsBookmarked: false,
			CallSign:     "PUMP-STATION",
			Color:        "#3fbf7f",
			Features: []string{
				"data-forwarding",
				"rate-limits",
				"releases",
			},
			Status:       "active",
			Organization: expectedOrganization,
		},
	}
	assert.Equal(t, expected, projects)
}

func TestProjectService_Get(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/0/projects/the-interstellar-jurisdiction/pump-station/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
    	"subjectPrefix": null,
    	"defaultEnvironment": null,
    	"features": [
    		"data-forwarding",
    		"rate-limits",
    		"releases"
    	],
    	"color": "#3fbf7f",
    	"platforms": [],
    	"plugins": [
    		{
    			"status": "unknown",
    			"assets": [],
    			"isTestable": true,
    			"contexts": [],
    			"doc": "",
    			"enabled": false,
    			"name": "WebHooks",
    			"canDisable": true,
    			"type": "notification",
    			"id": "webhooks",
    			"metadata": {}
    		}
    	],
    	"callSignReviewed": false,
    	"id": "2",
    	"digestsMinDelay": 300,
    	"firstEvent": null,
    	"digestsMaxDelay": 1800,
    	"processingIssues": 0,
    	"status": "active",
    	"isPublic": false,
    	"dateCreated": "2017-07-18T19:29:24.793Z",
    	"subjectTemplate": "[$project] ${tag:level}: $title",
    	"slug": "pump-station",
    	"name": "Pump Station",
    	"isBookmarked": false,
    	"callSign": "PUMP-STATION",
    	"team": {
    		"slug": "powerful-abolitionist",
    		"name": "Powerful Abolitionist",
    		"hasAccess": true,
    		"isPending": false,
    		"dateCreated": "2017-07-18T19:29:24.743Z",
    		"isMember": false,
    		"id": "2"
    	},
    	"organization": {
    		"name": "The Interstellar Jurisdiction",
    		"slug": "the-interstellar-jurisdiction",
    		"avatar": {
    			"avatarUuid": null,
    			"avatarType": "letter_avatar"
    		},
    		"dateCreated": "2017-07-18T19:29:24.565Z",
    		"id": "2",
    		"isEarlyAdopter": false
    	},
    	"options": {
    		"sentry:csp_ignored_sources_defaults": true,
    		"sentry:scrub_defaults": true,
    		"sentry:origins": "*",
    		"sentry:resolve_age": 0,
    		"sentry:sensitive_fields": [],
    		"sentry:scrub_data": true,
    		"sentry:reprocessing_active": false,
    		"sentry:csp_ignored_sources": "",
    		"filters:blacklisted_ips": "",
    		"sentry:safe_fields": [],
    		"feedback:branding": true,
    		"sentry:default_environment": null
    	}
    }`)
	})

	client := NewClient(httpClient, nil, "")
	project, _, err := client.Projects.Get("the-interstellar-jurisdiction", "pump-station")
	assert.NoError(t, err)
	expected := &Project{
		ID:           "2",
		Slug:         "pump-station",
		Name:         "Pump Station",
		DateCreated:  mustParseTime("2017-07-18T19:29:24.793Z"),
		IsPublic:     false,
		IsBookmarked: false,
		CallSign:     "PUMP-STATION",
		Color:        "#3fbf7f",
		Features: []string{
			"data-forwarding",
			"rate-limits",
			"releases",
		},
		Status: "active",
		Team: Team{
			ID:          "2",
			Slug:        "powerful-abolitionist",
			Name:        "Powerful Abolitionist",
			DateCreated: mustParseTime("2017-07-18T19:29:24.743Z"),
			HasAccess:   true,
			IsPending:   false,
			IsMember:    false,
		},
		Organization: Organization{
			ID:          "2",
			Slug:        "the-interstellar-jurisdiction",
			Name:        "The Interstellar Jurisdiction",
			DateCreated: mustParseTime("2017-07-18T19:29:24.565Z"),
			Avatar: OrganizationAvatar{
				Type: "letter_avatar",
			},
			IsEarlyAdopter: false,
		},
	}
	assert.Equal(t, expected, project)
}
