#include <iostream>
#include <map>
#include <string>
#include <algorithm>
using namespace std;

bool checkPalindromeAvailable(map<char, int>& alphabet){
	int oddNumber = 0;
	for(char ch = 'A'; ch <= 'Z'; ch++){
		if(alphabet[ch] % 2 != 0){
			oddNumber++;
		}
	}
	if(oddNumber > 1) return false;
	else return true;
}
string multiplyString(char ch, int multiplier){
	string result = "";
	for(int i=0; i<multiplier; i++){
		result += ch;
	}
	return result;
}
string getPalindrome(map<char, int>& alphabet){
	string halfStr; // will be reversed
	string middleChar = "";
	for(char ch = 'A'; ch <= 'Z'; ch++){
		if(alphabet[ch] > 0){
			if(alphabet[ch] % 2 != 0){
				middleChar += ch;
			}
			halfStr += multiplyString(ch, (alphabet[ch]/2));
		}
	}
	string halfTmpStr = halfStr;
	reverse(halfStr.begin(), halfStr.end());
	return (halfTmpStr + middleChar + halfStr);
}
int main(int argc, char* argv[]){
	map<char, int> alphabet;
	for(char ch = 'A'; ch <= 'Z'; ch++){
		alphabet[ch] = 0;
	}
	string input;
	cin >> input;
	for(int i=0; i<input.length(); i++){
		alphabet[input[i]]++;
	}
	if(!checkPalindromeAvailable(alphabet)){
		cout << "I'm Sorry Hansoo" << '\n';
	}
	else{
		string result = getPalindrome(alphabet);	
		cout << result << '\n';
	}
	return 0;
}
