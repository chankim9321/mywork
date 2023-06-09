#include <iostream>
#include <string>
using namespace std;
bool isVowel(char ch){
	if(ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u') return true;
	else return false;
}
bool ruleCheck(string input){
	int vowelCnt = 0; // 모음 카운팅
	int consonantCnt = 0; // 자음 카운팅
	bool vowelExist = false; // 모음 포함 여부
	char preChar;
	for(int i = 0; i<input.length(); i++){
		if(isVowel(input[i])){ // 모음 검사
			if(i>0){ // 같은 글자 연속성 체크
				if((preChar == 'e' && input[i] == 'e') || (preChar == 'o' && input[i] == 'o')){

				}
				else if(preChar == input[i]){
					return false;
				}
			}
			vowelCnt++;
			consonantCnt = 0;
			vowelExist = true;
		}
		else{ // 자음 검사
			if(i > 0 && (preChar == input[i])){
				return false;
			}
			consonantCnt++;
			vowelCnt = 0;
		}
		if(consonantCnt == 3 || vowelCnt == 3) return false;
		preChar = input[i];
	}
	if(vowelExist) return true;
	else return false;
}
int main(int argc, char* argv[]){
	ios_base::sync_with_stdio(false);
	cin.tie(0);
	cout.tie(0);

	while(1){
		string input;
		cin >> input;
		if(input == "end") break;
		if(ruleCheck(input)){
			cout << "<" << input << ">" << " is acceptable.\n";
		}
		else{
			cout << "<" << input << ">" << " is not acceptable.\n";
		}
	}
	return 0;
}
