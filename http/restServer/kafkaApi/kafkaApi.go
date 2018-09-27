package kafkaApi

import (
	"github.com/diadata-org/api-golang/dia/helpers/kafkaHelper"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

var (
	Apis = map[int]*RestApi{
		kafkaHelper.TopicFiltersBlock: NewRestApi(kafkaHelper.TopicFiltersBlock),
		kafkaHelper.TopicTrades:       NewRestApi(kafkaHelper.TopicTrades),
		kafkaHelper.TopicTradesBlock:  NewRestApi(kafkaHelper.TopicTradesBlock),
	}
)

// @hello
// returns some kafka messages
type resultApi struct {
	Offset   int64         `json:"offset"`
	Messages []interface{} `json:"messages"`
}

type RestApi struct {
	topic  int
	writer *kafka.Writer //TOFIX reorganise the api
}

func SendError(c *gin.Context, errorCode int, err error) {
	c.JSON(errorCode,
		&APIError{
			ErrorCode:    errorCode,
			ErrorMessage: err.Error(),
		})
}

func getOffset(c *gin.Context) int64 {
	offsetQuery := c.Query("offset")
	if offsetQuery == "" {
		offsetQuery = "-1" // will go to the latest element if no offset is specified
	}
	log.Printf("offsetQuery %v", offsetQuery)
	offset, _ := strconv.Atoi(offsetQuery)
	return int64(offset)
}

func Process(c *gin.Context, topic int) {
	elements, _ := strconv.Atoi(c.Query("elements"))
	result, err := Apis[topic].Get(getOffset(c), elements)
	if err == nil {
		c.JSON(http.StatusOK, result)
	} else {
		SendError(c, http.StatusInternalServerError, err)
	}
}

//
// @Summary Add a new pet to the store
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Param   some_id     path    int     true        "Some ID"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} restApi.APIError "We need ID!!"
// @Failure 404 {object} restApi.APIError "Can not find ID"
// @Router /testapi/get-string-by-int/{some_id} [get]

func (s *RestApi) Get(offset int64, elements int) (map[string]interface{}, error) {

	if (elements == 0) || (elements > 100) {
		elements = 100
	}

	result := &resultApi{}

	maxOffset, err := kafkaHelper.ReadOffset(s.topic)

	if err != nil {
		return nil, err
	}

	maxOffset--

	if offset > maxOffset {
		offset = maxOffset
	}
	if offset < 0 {
		offset = maxOffset
	}
	result.Offset = offset

	nbElements := int(maxOffset - offset + 1)
	if nbElements > elements {
		nbElements = elements
	}
	log.Printf("Get: maxOffset %v offset:%v nbElements:%v ", maxOffset, offset, nbElements)

	element, err := kafkaHelper.GetElements(s.topic, offset, nbElements)

	if err != nil {
		return nil, err
	}
	result.Messages = append(result.Messages, element)

	r := map[string]interface{}{
		"Result": result,
	}

	return r, nil

}

func NewRestApi(topic int) *RestApi {
	s := &RestApi{
		topic: topic,
	}
	return s
}
