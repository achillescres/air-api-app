package parser

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	config2 "github.com/achillescres/saina-api/internal/config"
	"github.com/achillescres/saina-api/internal/domain/service"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"strings"
)

type TaisParser interface {
	ParseFirstTaisFile(ctx context.Context) (map[int]error, error)
}

type taisParser struct {
	parserService service.ParserService
	cfg           config2.TaisParserConfig
}

var _ TaisParser = (*taisParser)(nil)

func NewTaisParser(parserService service.ParserService, cfg config2.TaisParserConfig) TaisParser {
	return &taisParser{parserService: parserService, cfg: cfg}
}

func (tP *taisParser) ParseFirstTaisFile(ctx context.Context) (map[int]error, error) {
	env := config2.Env()
	taisDirPath := path.Join(env.ProjectAbsPath, tP.cfg.TaisDirPath)
	inDir, err := os.ReadDir(taisDirPath)
	if err != nil {
		// TODO remove fatal drop, add logic to save the system from deprecated data(outer layer work)
		log.Fatalf("(WillRemLog)fatal scanning tais directory=%s: %s\n", taisDirPath, err.Error())
		return map[int]error{}, err
	}

	taisFileName := ""
	for _, entry := range inDir {
		if !entry.IsDir() {
			taisFileName = entry.Name()
		}
	}

	if taisFileName == "" {
		log.Fatalf("(WillRemLog)fatal didnt find tais file in tais dir=%s\n", taisDirPath) // TODO remove fatals
		return map[int]error{}, errors.New(fmt.Sprintf("error didnt find tais file in tais dir=%s\n", taisDirPath))
	}

	taisFilePath := path.Join(taisDirPath, taisFileName)
	f, err := os.OpenFile(taisFilePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("(WillRemLog)error opening file for parse: %s\n", err.Error()) // TODO remove fatals
		return map[int]error{}, err
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			// TODO add normal logic to prevent memory leaks from unclosed files
			log.Errorf("error closing tais file=%s: %s\n", f.Name(), err.Error())
			f.Close() // TODO this is not best practice i'd say
		}
	}(f)

	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanLines)
	if !sc.Scan() {
		log.Errorln("error tais parse file is empty")
		return map[int]error{}, errors.New("error parse file is empty")
	} else {
		sc.Text() // Meta line
	}

	rows := make([]string, 0, 1800)
	for sc.Scan() {
		rows = append(rows, sc.Text())
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("(WillRemLog)error reading tais file for parse: %s\n", err.Error()) // TODO remove fatals
		return map[int]error{}, err
	}

	errs := map[int]error{}
	for i, row := range rows {
		fields := strings.Fields(strings.TrimSpace(row))
		err := tP.parserService.ParseFields(ctx, fields)
		if err != nil {
			errs[i] = err
			log.Errorf("error parsing row of tais file: %s", err.Error())
		}
	}

	return errs, nil
}
