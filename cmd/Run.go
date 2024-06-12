package main

import (
    "log"
    "os/exec"
    "sync"
)

var wg sync.WaitGroup

func main() {
    // Путь к вашим .exe файлам
    exePathtoWebApp := `C:\Users\VladislavSCV\OneDrive\Desktop\VSC\GO\GoCode\cmd\web.exe`
    exePathtoGRPCService := `C:\Users\VladislavSCV\OneDrive\Desktop\VSC\GO\GoCode\cmd\grpc.exe`

    listOfPaths := []string{exePathtoWebApp, exePathtoGRPCService}

    // Добавляем к счетчику количество горутин
    wg.Add(len(listOfPaths))

    for _, path := range listOfPaths {
        go execFile(path)
    }

    wg.Wait()
}

func execFile(path string) {
    defer wg.Done()
    
    cmd := exec.Command(path)

    // Запуск команды и проверка ошибок
    output, err := cmd.CombinedOutput()
    if err != nil {
        log.Printf("Ошибка при запуске .exe файла в файле: %s ошибка: %v\nВывод: %s", path, err, output)
        return
    }

    log.Printf("Программа %s успешно запущена\nВывод: %s", path, output)
}