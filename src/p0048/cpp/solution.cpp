class Solution {
public:
    void rotate(vector<vector<int>>& matrix) {
       auto n = matrix.size();
       int i = 0, j = n - 1;
       while (i < j) {
           swap(matrix[i++], matrix[j--]);
       }
       for (int i = 0; i < matrix.size(); i++) {
           for (int j = i; j < matrix[i].size(); j++) {
               swap(matrix[i][j], matrix[j][i]);
           }
       }
    }
};