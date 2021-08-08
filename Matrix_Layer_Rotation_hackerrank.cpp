#include <bits/stdc++.h>
#include <iostream>
using namespace std;

string ltrim(const string &);
string rtrim(const string &);
vector<string> split(const string &);

// Complete the matrixRotation function below.
void matrixRotation(vector<vector<int>> matrix, int r) {
 int const matrixh=matrix.size();
int  const  matrixv=matrix[0].size();
int const circ=(matrixh+matrixv-2)*2;
int beginh;
int endh;
int beginv;
int endv;
int prov[4];
int store[4];
int loopv;
int mod;
vector<int> ring;
int const sv=(matrixh>matrixv)?matrixv:matrixh;
for (int rin=0;rin<sv/2;++rin){
    mod=r%((matrixh+matrixv-2*rin-2*(rin+1))*2);
    ring.push_back(mod);
}

for (int incre=0;incre<circ;++incre){
   loopv = (matrixh>matrixv)?matrixh:matrixv;
   ++loopv;
    beginh=0;
    endh=matrixh-1;
    beginv=0;
    endv=matrixv-1;
for (int s=0;s<sv/2;++s)
{
    loopv-=1;
     beginh=s;
     endh=matrixh-1-s;
     beginv=s;
     endv=matrixv-1-s;
    store[0]=matrix[s][beginv];
    store[1]= matrix[endh][endv];
    store[2]=matrix[endh][s];
    store[3]=matrix[beginh][endv];
    if (incre>=ring[s])
        continue;
for (int i=1;i<loopv;++i){
    copy(begin(store),end(store), begin(prov));
    if (endh-i>=s){
    store[0]=matrix[s+i][beginv];
    matrix[s+i][beginv]=prov[0];
    store[1]= matrix[endh-i][endv];
    matrix[endh-i][endv]=prov[1];
}
     if(endv-i>=s){
    store[2]=matrix[endh][s+i];
    matrix[endh][s+i]=prov[2];
    store[3]=matrix[beginh][endv-i];
    matrix[beginh][endv-i]=prov[3];
}
}
}

}
for(auto const &ele:matrix){
    for(auto const &elem:ele)
       cout<<(elem)<<" ";
    cout<<("\n");
}
}

int main()
{
    string mnr_temp;
    getline(cin, mnr_temp);

    vector<string> mnr = split(rtrim(mnr_temp));

    int m = stoi(mnr[0]);

    int n = stoi(mnr[1]);

    int r = stoi(mnr[2]);

    vector<vector<int>> matrix(m);

    for (int i = 0; i < m; i++) {
        matrix[i].resize(n);

        string matrix_row_temp_temp;
        getline(cin, matrix_row_temp_temp);

        vector<string> matrix_row_temp = split(rtrim(matrix_row_temp_temp));

        for (int j = 0; j < n; j++) {
            int matrix_row_item = stoi(matrix_row_temp[j]);

            matrix[i][j] = matrix_row_item;
        }
    }

    matrixRotation(matrix, r);

    return 0;
}

string ltrim(const string &str) {
    string s(str);

    s.erase(
        s.begin(),
        find_if(s.begin(), s.end(), not1(ptr_fun<int, int>(isspace)))
    );

    return s;
}

string rtrim(const string &str) {
    string s(str);

    s.erase(
        find_if(s.rbegin(), s.rend(), not1(ptr_fun<int, int>(isspace))).base(),
        s.end()
    );

    return s;
}

vector<string> split(const string &str) {
    vector<string> tokens;

    string::size_type start = 0;
    string::size_type end = 0;

    while ((end = str.find(" ", start)) != string::npos) {
        tokens.push_back(str.substr(start, end - start));

        start = end + 1;
    }

    tokens.push_back(str.substr(start));

    return tokens;
}
