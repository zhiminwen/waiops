package waiops

import (
	"encoding/json"
	"fmt"
	"reflect"
	"slices"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/dsnet/try"
	"github.com/zhiminwen/quote"
)

type EvLink struct {
	LinkType    string `json:"linkType"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
}

type EvResource struct {
	Name         string `json:"name"`
	SourceId     string `json:"sourceId"`
	Hostname     string `json:"hostname"`
	IpAddress    string `json:"ipAddress"`
	Service      string `json:"service"`
	Port         int    `json:"port"`
	Interface    string `json:"interface"`
	Application  string `json:"application"`
	Controller   string `json:"controller"`
	Component    string `json:"component"`
	Cluster      string `json:"cluster"`
	Location     string `json:"location"`
	AccessScope  string `json:"accessScope"`
	ConnectionId string `json:"connectionId"`
	ScopeId      string `json:"scopeId"`
}

type EvType struct {
	Classification string `json:"classification"`
	EventType      string `json:"eventType"`
	Condition      string `json:"condition"`
}

type EvTime time.Time

func (t EvTime) MarshalJSON() ([]byte, error) {
	// 2023-08-23T20:41:12.420Z
	return json.Marshal(time.Time(t).Format("2006-01-02T15:04:05.000Z"))
}

func (t *EvTime) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	parsedTime, err := time.Parse("2006-01-02T15:04:05.000Z", s)
	if err != nil {
		return err
	}
	*t = EvTime(parsedTime)
	return nil
}

type EvEvent struct {
	Id             string     `json:"id"`
	OccurrenceTime EvTime     `json:"occurrenceTime"`
	Summary        string     `json:"summary"`
	Severity       int        `json:"severity"`
	Sender         EvResource `json:"sender"`
	Resource       EvResource `json:"resource"`
	ExpirySeconds  int        `json:"expirySeconds"`
	Links          []EvLink   `json:"links"`
	Type           EvType     `json:"type"`
}

type EvAlert struct {
	Id                  string     `json:"id"`
	State               string     `json:"state"`
	EventCount          int        `json:"eventCount"`
	Acknowledged        bool       `json:"acknowledged"`
	Team                string     `json:"team"`
	Owner               string     `json:"owner"`
	DeduplicationKey    string     `json:"deduplicationKey"`
	Signature           string     `json:"signature"`
	OccurrenceTime      EvTime     `json:"occurrenceTime"`
	FirstOccurrenceTime EvTime     `json:"firstOccurrenceTime"`
	LastOccurrenceTime  EvTime     `json:"lastOccurrenceTime"`
	Summary             string     `json:"summary"`
	LangId              string     `json:"langId"`
	Severity            int        `json:"severity"`
	Sender              EvResource `json:"sender"`
	Resource            EvResource `json:"resource"`
	Type                EvType     `json:"type"`
	ExpirySeconds       int        `json:"expirySeconds"`
	Links               []EvLink   `json:"links"`

	// Details             map[string]string `json:"details"`
	// Insight string `json:"insight"`
}

type EvChangeNotification struct {
	TentantId        string  `json:"tenantid"`
	RequestId        string  `json:"requestid"`
	NotificationTime EvTime  `json:"notificationTime"`
	Type             string  `json:"type"`
	EntityType       string  `json:"entityType"`
	Entity           EvAlert `json:"entity"`
}

func randomResource() EvResource {
	return EvResource{
		Name:      gofakeit.AppName(),
		SourceId:  gofakeit.UUID(),
		Hostname:  gofakeit.DomainName(),
		IpAddress: gofakeit.IPv4Address(),
		Service:   gofakeit.Name(),
		Port:      gofakeit.Number(20, 65535),
		Interface: gofakeit.RandomString(quote.Line(`
			eth0
			eth1
			eth2
			wlan0
			wlan1
			wlan2
			lo
			eno1
			eno2
			eno3
		`)),
		Application: gofakeit.AppName(),
		Controller:  gofakeit.Noun(),
		Component: gofakeit.RandomString(quote.Line(`
			Database management systems (DBMS)
			Relational databases (RDBMS)
			NoSQL databases
			Data warehouses
			Data lakes
			HTTP servers
			Application servers
			Messaging queues
			API gateways
			Load balancers
			Authentication servers
			Authorization servers
			Encryption tools
			Caching engines
			Search engines
			Logging and monitoring tools
			MySQL
			PostgreSQL
			Oracle
			MongoDB
			Cassandra
			Redis
			Apache HTTP Server
			Nginx
			Microsoft IIS
			Apache Tomcat
			IBM WebSphere
			Oracle WebLogic
			Apache Kafka
			RabbitMQ
			Amazon SQS
			NGINX
			Amazon API Gateway
			Google Cloud Endpoints
			HAProxy
			LDAP
			Active Directory
			Kerberos
			OAuth
			OpenID Connect
			RBAC
			SSL/TLS
			PGP
			AES
			Elasticsearch
			Apache Solr
			Google Search Appliance
			Log4j
			Logstash
			Nagios
		`)),
		Cluster:      gofakeit.Word(),
		Location:     gofakeit.City(),
		AccessScope:  gofakeit.Word(),
		ConnectionId: gofakeit.UUID(),
		ScopeId:      gofakeit.UUID(),
	}
}

func randomLinks() []EvLink {
	links := []EvLink{}
	for i := 0; i < gofakeit.Number(1, 2); i++ {
		links = append(links, EvLink{
			LinkType:    gofakeit.Word(),
			Name:        gofakeit.Noun(),
			Description: gofakeit.Sentence(20),
			Url:         gofakeit.URL(),
		})
	}
	return links
}

func randomType() EvType {
	return EvType{
		Classification: gofakeit.RandomString(quote.Line(`
			System status
			Threshold breach
			Utilization
			Performance metrics
			Uptime
			Downtime
			Latency
			Throughput
			Response time
			Error rate
		`)),
		// EventType: "problem",
		EventType: gofakeit.RandomString([]string{"problem", "resolution"}),
		Condition: gofakeit.HackerAdjective() + " " + gofakeit.HackerNoun(),
	}
}

func NewRandomEvent() EvEvent {
	return EvEvent{
		Id:             gofakeit.UUID(),
		OccurrenceTime: EvTime(gofakeit.DateRange(time.Now().AddDate(0, 0, -1), time.Now())),
		Summary:        gofakeit.Sentence(50),
		Severity:       gofakeit.Number(1, 6),
		Sender:         randomResource(),
		Resource:       randomResource(),

		ExpirySeconds: gofakeit.Number(300, 1000),
		Links:         randomLinks(),
		Type:          randomType(),
	}
}

func NewRandomAlert() EvAlert {
	alert := EvAlert{
		Id:    gofakeit.UUID(),
		State: "open",
		// State:        gofakeit.RandomString([]string{"open", "clear", "closed"}),
		EventCount:   gofakeit.Number(1, 10),
		Acknowledged: gofakeit.Bool(),
		Team:         gofakeit.NounCollectivePeople(),
		Owner:        gofakeit.Name(),

		FirstOccurrenceTime: EvTime(gofakeit.DateRange(time.Now().AddDate(0, 0, -7), time.Now())),
		Summary:             gofakeit.HackerPhrase(), //Sentence(50),
		LangId:              gofakeit.RandomString([]string{"eng", "fra", "deu", "jpn", "kor", "zho"}),
		Severity:            gofakeit.Number(1, 6),
		Sender:              randomResource(),
		Type:                randomType(),
		ExpirySeconds:       gofakeit.Number(0, 3000),
		Links:               randomLinks(),
	}
	alert.OccurrenceTime = alert.FirstOccurrenceTime
	alert.LastOccurrenceTime = EvTime(gofakeit.DateRange(time.Time(alert.FirstOccurrenceTime), time.Now()))

	alert.SetResource(randomResource()) //also set DeduplicationKey and Signature by resource value
	return alert
}

func (a *EvAlert) SetEventType(classification, pOrr, cond string) *EvAlert {
	a.Type = EvType{
		Classification: classification,
		EventType:      pOrr,
		Condition:      cond,
	}
	a.UpdateDedupKeyAndSignature()
	return a
}

func (a *EvAlert) SetEventTypeAsProblemOrResolution(pOrr string) *EvAlert {
	a.Type.EventType = pOrr
	a.UpdateDedupKeyAndSignature()
	return a
}

// Set Resource, and set DeduplicationKey, Signature based on Resource and type classification
func (a *EvAlert) SetResource(res EvResource) *EvAlert {
	a.Resource = res
	a.UpdateDedupKeyAndSignature()
	return a
}

func (a *EvAlert) UpdateDedupKeyAndSignature() *EvAlert {
	dkey := map[string]string{}

	ref := reflect.ValueOf(a.Resource)
	keys := []string{}
	for i := 0; i < ref.NumField(); i++ {
		k := ref.Type().Field(i).Tag.Get("json")
		f := ref.Field(i)
		val := ref.Field(i).String()

		switch f.Kind() {
		case reflect.Int:
			val = fmt.Sprintf("%d", f.Int())
		}
		if val != "" && val != "0" {
			dkey[k] = val
			keys = append(keys, k)
		}
	}

	slices.Sort(keys)
	signatures := []string{}
	for _, k := range keys {
		signatures = append(signatures, fmt.Sprintf("%s=%s", k, dkey[k]))
	}
	sig := fmt.Sprintf("{%s}-%s-%s", strings.Join(signatures, ","), a.Type.Classification, a.Type.Condition)
	a.DeduplicationKey = sig
	a.Signature = sig

	return a
}

func (a *EvAlert) SetOccurrenceTime(first, last time.Time, count int) *EvAlert {
	a.OccurrenceTime = EvTime(first)
	a.FirstOccurrenceTime = EvTime(first)
	a.LastOccurrenceTime = EvTime(last)
	a.EventCount = count
	return a
}

func (a *EvAlert) SetExpiration(seconds int) *EvAlert {
	a.ExpirySeconds = seconds
	return a
}

func (a *EvAlert) AsJson() []byte {
	payload := try.E1(json.MarshalIndent(a, "", "  "))
	return payload

}
