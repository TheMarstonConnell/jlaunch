package core

import (
	"fmt"
	"github.com/JackalLabs/jackalgo/handlers/file_io_handler"
	"github.com/JackalLabs/jackalgo/handlers/file_upload_handler"
	"github.com/JackalLabs/jackalgo/handlers/folder_handler"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
)

func DeleteFolder(path string) error {
	wallet, fileIo, err := InitWalletSession()
	if err != nil {
		return fmt.Errorf("cannot save %s | %w", path, err)
	}

	_ = wallet

	return Delete(fileIo, path)

	return nil
}

func Delete(fileIo *file_io_handler.FileIoHandler, path string) error {
	folder := filepath.Base(path)
	parent := filepath.Dir(path)
	fmt.Println(parent)
	parentFolder, msgs, err := fileIo.LoadNestedFolder(parent)
	if err != nil {
		return fmt.Errorf("cannot load folder from chain %s | %w", path, err)
	}
	fmt.Println(parentFolder.GetWhereAmI())
	if len(msgs) > 0 {
		err = fileIo.SignAndBroadcast(msgs)
		if err != nil {
			return fmt.Errorf("failed to create folder %s on chain | %w", path, err)
		}
	}

	err = fileIo.DeleteTargets([]string{folder}, parentFolder)
	if err != nil {
		return fmt.Errorf("cannot load folder from chain %s | %w", path, err)
	}

	return nil
}

func SaveFolder(folder string, path string) error {
	wallet, fileIo, err := InitWalletSession()
	if err != nil {
		return fmt.Errorf("cannot save %s | %w", folder, err)
	}

	_ = wallet

	return Save(fileIo, filepath.Dir(folder), filepath.Base(folder), path)

	return nil
}

func SaveFile(file string, path string) error {
	wallet, fileIo, err := InitWalletSession()
	if err != nil {
		return fmt.Errorf("cannot save %s due to wallet/fileio issues | %w", file, err)
	}

	_ = wallet

	folder, err := fileIo.DownloadFolder(path)
	if err != nil {
		return fmt.Errorf("could not get folder %s | %w", path, err)
	}

	return saveFile(filepath.Base(file), filepath.Dir(file), fileIo, folder)

	return nil
}

func Save(fileIo *file_io_handler.FileIoHandler, root string, dir string, path string) error {

	l := filepath.Join(root, dir)

	log.Info().Msgf("%s / %s", root, dir)
	log.Info().Msgf("%s / %s", path, dir)

	newPath := filepath.Join(path, filepath.Base(dir))

	parentFolder, msgs, err := fileIo.LoadNestedFolder(newPath)
	if err != nil {
		return fmt.Errorf("cannot load folder from chain %s | %w", l, err)
	}

	if len(msgs) > 0 {
		err = fileIo.SignAndBroadcast(msgs)
		if err != nil {
			return fmt.Errorf("failed to create folder %s on chain | %w", l, err)
		}
	}

	entries, err := os.ReadDir(l)
	if err != nil {
		return fmt.Errorf("cannot read folder entries %s | %w", l, err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			if entry.Name()[:1] == "." {
				continue
			}

			err := Save(fileIo, root, filepath.Join(dir, entry.Name()), newPath)
			if err != nil {
				log.Warn().Err(err)
				continue
			}
		}

		err := saveFile(entry.Name(), filepath.Join(root, dir), fileIo, parentFolder)
		if err != nil {
			log.Warn().Err(err)
			continue
		}
	}

	return nil
}

func saveFile(fileName string, filePath string, fileIo *file_io_handler.FileIoHandler, parent *folder_handler.FolderHandler) error {

	if fileName[:1] == "." {
		return nil
	}

	p := filepath.Join(filePath, fileName)

	log.Info().Msgf("saving %s to %s", p, parent.GetMyPath())

	data, err := os.ReadFile(p)
	if err != nil {
		return fmt.Errorf("cannot find file %s on disk | %w", p, err)
	}

	upload, err := file_upload_handler.TrackVirtualFile(data, fileName, parent.GetMyPath())
	if err != nil {
		return fmt.Errorf("cannot track file %s | %w", fileName, err)
	}

	_, _, _, err = fileIo.StaggeredUploadFiles([]*file_upload_handler.FileUploadHandler{upload}, parent, true)
	if err != nil {
		return fmt.Errorf("cannot save file %s | %w", fileName, err)
	}

	return nil
}
