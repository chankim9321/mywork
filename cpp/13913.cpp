#include <iostream>
#include <queue>
#include <vector>
#include <algorithm>
using namespace std;

int cnt[100001];
int next_x[3] ={-1,1,2};
bool visitedPos[100001] = { false, };
vector<int> road(100000);

void sol(int subinPos, int brotherPos){
	queue<pair<pair<int, int>, vector<int>>> q;
	vector<int> v;
	q.push({{subinPos, 0}, v});
	while(!q.empty()){
		pair<int, int> present = q.front().first;
		vector<int> visited = q.front().second;
		q.pop();
		cnt[present.first] = min(cnt[present.first], present.second);
		visited.push_back(present.first);
		visitedPos[present.first] = true;
		if(present.first == brotherPos){ // 현재 위치가 동생이 있는 위치라면 지나온 길을 가장 짧은 것으로 업데이트
			if(road.size() > visited.size()){
				road = visited;
			}
		}
		for(int i=0; i<3; i++){
			int nextPos;
			if(i==2){
				nextPos = present.first * next_x[i];
			}
			else{
				nextPos = present.first + next_x[i];
			}
			if(nextPos < 100001 && !visitedPos[nextPos]){
				q.push({{nextPos, present.second+1}, visited});
			}
		}
	}
}

int main(int argc, char* argv[]){
	ios_base::sync_with_stdio(false); 
	cin.tie(0);
	cout.tie(0);
	fill_n(&cnt[0], 100001, 100001);
	int subinPos, brotherPos;
	cin >> subinPos >> brotherPos;
	if(subinPos > brotherPos){ // 수빈의 포지션이 동생보다 앞선다면 뒤로 한칸씩 가는 수밖에 없음.
		cout << subinPos - brotherPos << '\n';
		for(int i=subinPos; i>=brotherPos; i--){
			cout << i << " ";
		}
		cout << '\n';
		return 0;
	}
	else {
		sol(subinPos, brotherPos);
		cout << cnt[brotherPos] << '\n';
		for(int i=0; i<road.size(); i++){
			cout << road[i] << " ";
		}
		cout << '\n';
		return 0;
	}
}