package scraper

import (
	"context"
	"time"

	"github.com/chromedp/chromedp"
)

// Определяем тип Result для хранения данных о карте и сумме
type Result struct {
	CardNumber string
	Amount     string
}

func ScrapeData(ctx context.Context, source string) ([]Result, error) {
	var results []Result
	var err error

	switch source {
	case "bcc":
		results, err = ScrapeDataBCC(ctx)
	case "halyk":
		results, err = ScrapeDataHalyk(ctx)
	case "kaspi":
		results, err = ScrapeDataKaspi(ctx)
	case "jusan":
		results, err = ScrapeDataJusan(ctx)
	case "bereke":
		results, err = ScrapeDataBereke(ctx)
	}

	if err != nil {
		return nil, err
	}

	return results, nil
}

// Функция для парсинга данных BCC
func ScrapeDataBCC(ctx context.Context) ([]Result, error) {
	var results []Result

	// Подключаемся к удалённому отладчику
	ctx, cancel := chromedp.NewRemoteAllocator(ctx, "http://localhost:9222")
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Navigate("https://my.prod.platcore.io/pay-out?limit=100&status=new"),
		chromedp.WaitVisible(`div.css-1q1bux1`, chromedp.ByQuery),
		chromedp.Sleep(5*time.Second),
		chromedp.Evaluate(`
            Array.from(document.querySelectorAll('tr.css-115hxry'))
                .map(row => {
                    const cardNumberElements = Array.from(row.querySelectorAll('div.css-1q1bux1'));
                    const cardNumberElement = cardNumberElements.find(el => 
                        ['446375', '462818', '489993', '526988', '539674', '400289', 
 						'490553', '490449', '490453', '404242', '404243', '404245', 
 						'444499', '517792', '525752', '536685', '403259', '429439', 
 						'526994', '521700', '533642', '530496', '516873', '418973', 
 						'440125', '404932', '423300', '423306', '532456', '444077', 
 						'441328', '516949'].some(prefix => el.innerText.trim().startsWith(prefix))
                    );

                    const amountElements = Array.from(row.querySelectorAll('div.chakra-stack.css-zdx2uo > div.css-1q1bux1'));
                    const amountElement = amountElements.find(el => el.innerText.trim().startsWith('-'));

                    return {
                        CardNumber: cardNumberElement ? cardNumberElement.innerText.trim() : "",
                        Amount: amountElement ? amountElement.innerText.trim() : ""
                    };
                })
                .filter(result => result.CardNumber !== "" && result.Amount !== "") 
        `, &results),
	)
	if len(results) == 0 {
		return results, nil
	}

	if err != nil {
		return nil, err
	}

	return results, nil
}

// Функция для парсинга данных Halyk
func ScrapeDataHalyk(ctx context.Context) ([]Result, error) {
	var results []Result

	// Подключаемся к удалённому отладчику
	ctx, cancel := chromedp.NewRemoteAllocator(ctx, "http://localhost:9222")
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Navigate("https://my.prod.platcore.io/pay-out?limit=100&status=new"),
		chromedp.WaitVisible(`div.css-1q1bux1`, chromedp.ByQuery),
		chromedp.Sleep(5*time.Second),
		chromedp.Evaluate(`
            Array.from(document.querySelectorAll('tr.css-115hxry'))
                .map(row => {
                    const cardNumberElements = Array.from(row.querySelectorAll('div.css-1q1bux1'));
                    const cardNumberElement = cardNumberElements.find(el => 
                        ['400303', '427704', '422126', '444482', '490472', '552204', 
 						'535451', '548319', '440563', '547089', '517511'].some(prefix => el.innerText.trim().startsWith(prefix))
                    );

                    const amountElements = Array.from(row.querySelectorAll('div.chakra-stack.css-zdx2uo > div.css-1q1bux1'));
                    const amountElement = amountElements.find(el => el.innerText.trim().startsWith('-'));

                    return {
                        CardNumber: cardNumberElement ? cardNumberElement.innerText.trim() : "",
                        Amount: amountElement ? amountElement.innerText.trim() : ""
                    };
                })
                .filter(result => result.CardNumber !== "" && result.Amount !== "") 
        `, &results),
	)
	if len(results) == 0 {
		return results, nil
	}

	if err != nil {
		return nil, err
	}

	return results, nil
}

// Функция для парсинга данных Kaspi
func ScrapeDataKaspi(ctx context.Context) ([]Result, error) {
	var results []Result

	// Подключаемся к удалённому отладчику
	ctx, cancel := chromedp.NewRemoteAllocator(ctx, "http://localhost:9222")
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Navigate("https://my.prod.platcore.io/pay-out?limit=100&status=new"),
		chromedp.WaitVisible(`div.css-1q1bux1`, chromedp.ByQuery),
		chromedp.Sleep(5*time.Second),
		chromedp.Evaluate(`
            Array.from(document.querySelectorAll('tr.css-115hxry'))
                .map(row => {
                    const cardNumberElements = Array.from(row.querySelectorAll('div.css-1q1bux1'));
                    const cardNumberElement = cardNumberElements.find(el => 
                        ['440043'].some(prefix => el.innerText.trim().startsWith(prefix))
                    );

                    const amountElements = Array.from(row.querySelectorAll('div.chakra-stack.css-zdx2uo > div.css-1q1bux1'));
                    const amountElement = amountElements.find(el => el.innerText.trim().startsWith('-'));

                    return {
                        CardNumber: cardNumberElement ? cardNumberElement.innerText.trim() : "",
                        Amount: amountElement ? amountElement.innerText.trim() : ""
                    };
                })
                .filter(result => result.CardNumber !== "" && result.Amount !== "") 
        `, &results),
	)
	if len(results) == 0 {
		return results, nil
	}

	if err != nil {
		return nil, err
	}

	return results, nil
}

// Функция для парсинга данных Jusan
func ScrapeDataJusan(ctx context.Context) ([]Result, error) {
	var results []Result

	// Подключаемся к удалённому отладчику
	ctx, cancel := chromedp.NewRemoteAllocator(ctx, "http://localhost:9222")
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Navigate("https://my.prod.platcore.io/pay-out?limit=100&status=new"),
		chromedp.WaitVisible(`div.css-1q1bux1`, chromedp.ByQuery),
		chromedp.Sleep(5*time.Second),
		chromedp.Evaluate(`
            Array.from(document.querySelectorAll('tr.css-115hxry'))
                .map(row => {
                    const cardNumberElements = Array.from(row.querySelectorAll('div.css-1q1bux1'));
                    const cardNumberElement = cardNumberElements.find(el => 
                        ['413264','458260','519170','526572','535650','539545'].some(prefix => el.innerText.trim().startsWith(prefix))
                    );

                    const amountElements = Array.from(row.querySelectorAll('div.chakra-stack.css-zdx2uo > div.css-1q1bux1'));
                    const amountElement = amountElements.find(el => el.innerText.trim().startsWith('-'));

                    return {
                        CardNumber: cardNumberElement ? cardNumberElement.innerText.trim() : "",
                        Amount: amountElement ? amountElement.innerText.trim() : ""
                    };
                })
                .filter(result => result.CardNumber !== "" && result.Amount !== "") 
        `, &results),
	)
	if len(results) == 0 {
		return results, nil
	}

	if err != nil {
		return nil, err
	}

	return results, nil
}

// Функция для парсинга данных Bereke
func ScrapeDataBereke(ctx context.Context) ([]Result, error) {
	var results []Result

	// Подключаемся к удалённому отладчику
	ctx, cancel := chromedp.NewRemoteAllocator(ctx, "http://localhost:9222")
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Navigate("https://my.prod.platcore.io/pay-out?limit=100&status=new"),
		chromedp.WaitVisible(`div.css-1q1bux1`, chromedp.ByQuery),
		chromedp.Sleep(5*time.Second),
		chromedp.Evaluate(`
            Array.from(document.querySelectorAll('tr.css-115hxry'))
                .map(row => {
                    const cardNumberElements = Array.from(row.querySelectorAll('div.css-1q1bux1'));
                    const cardNumberElement = cardNumberElements.find(el => 
                        ['457832','426343','542999','440256'].some(prefix => el.innerText.trim().startsWith(prefix))
                    );

                    const amountElements = Array.from(row.querySelectorAll('div.chakra-stack.css-zdx2uo > div.css-1q1bux1'));
                    const amountElement = amountElements.find(el => el.innerText.trim().startsWith('-'));

                    return {
                        CardNumber: cardNumberElement ? cardNumberElement.innerText.trim() : "",
                        Amount: amountElement ? amountElement.innerText.trim() : ""
                    };
                })
                .filter(result => result.CardNumber !== "" && result.Amount !== "") 
        `, &results),
	)
	if len(results) == 0 {
		return results, nil
	}
	if err != nil {
		return nil, err
	}

	return results, nil
}
