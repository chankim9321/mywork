#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

int tc;
int n,m,w;
bool BellmanFord(vector<vector<pair<int, int>>>& edge, int start){
	vector<int> dist(n+1, 10e8);
	dist[start] = 0;
	for(int i=1; i<=n; i++){ // 모든 노드
		int current = i;
		for(int j=0; j<edge[current].size(); j++){ // 모든 간선
			int next = edge[current][j].first;
			int weight = edge[current][j].second;
			if(dist[next] > dist[current] + weight){
				dist[next] = dist[current] + weight;
			}
		}
	}
	for(int i=1; i<=n; i++){ // 모든 노드
		int current = i;
		for(int j=0; j<edge[current].size(); j++){ // 모든 간선
			int next = edge[current][j].first;
			int weight = edge[current][j].second;
			if(dist[next] > dist[current] + weight) {
				return true;
			}
		}
	}
	return false;
}
void sol(){
	cin >> tc;
	for(int i=0; i<tc; i++){
		vector<vector<pair<int, int>>> edge;
		cin >> n >> m >> w;
		edge.resize(n+1);
		for(int j=0; j<m; j++){
			int from, to, weight;
			cin >> from >> to >> weight;
			edge[from].push_back({to, weight});
			edge[to].push_back({from, weight});
		}
		for(int j=0; j<w; j++){
			int from, to, weight;
			cin >> from >> to >> weight;
			edge[from].push_back({to, -1 * weight});
		}
		if(BellmanFord(edge, 1)) cout << "YES" << '\n'; // 음수 사이클이 존재한다면
		else cout << "NO" << '\n';
	}
}
int main(int argc, char* argv[]){
	ios_base::sync_with_stdio(false);
	cin.tie(0);
	cout.tie(0);

	sol();
	return 0;
}
