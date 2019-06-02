import UIKit

struct Diary: Codable {
    let id: Int
    let sender: String
    let title: String
    let content: String
}

class ViewController: UIViewController {

    override func viewDidLoad() {
        super.viewDidLoad()
        let url = URL(string: "http://localhost:5000/api/diary/findall")!
        URLSession.shared.dataTask(with: url) { (data, response, error) in
            if let error = error {
                fatalError("\(error)")
            }
            if let data = data, let diaries = try? JSONDecoder().decode([Diary].self, from: data) {
                for diary in diaries {
                    print(diary)
                }
            }
        }.resume()
    }


}

