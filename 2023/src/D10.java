import java.util.ArrayList;
import java.util.List;

import utils.InputFile;

class Pipe{
    char type;
    int[] pos;
    Pipe prev;
    Pipe next;
    boolean reachableFromOutsideLoop;
    Pipe(char type, int x, int y){
        this.type = type;
        this.pos = new int[]{x, y};
        this.prev = null;
        this.next = null;
        this.reachableFromOutsideLoop = false;
    }
    public String toString() {
        return this.type+" ";
    }
}

public class D10 {
    static List<String> raw;
    static Pipe[][] pipes;
    static Pipe start;
    static void parseInput(){
        raw = InputFile.readFile("./Tim3/src/input/10.txt");
        //raw = InputFile.readFile("./Tim3/src/input/10Sample.txt");
        //raw = InputFile.readFile("./Tim3/src/input/10SampleP2.txt");
        int maxX = raw.get(0).length();
        int maxY = raw.size();
        pipes = new Pipe[maxY][maxX];
        for (int y=0; y<maxY; y++){
            for (int x=0; x<maxX; x++){
                char type = raw.get(y).charAt(x);
                Pipe pipe = new Pipe(type, x, y);
                pipes[y][x] = pipe;
                if (type == 'S') start = pipe;
            }
        }
        for (int y=0; y<maxY; y++){
            for (int x=0; x<maxX; x++){
                Pipe pipe = pipes[y][x];
                int prevX;
                int prevY;
                int nextX;
                int nextY;
                switch (pipe.type) {
                    case '|':
                        prevX = x;
                        prevY = y-1;
                        nextX = x;
                        nextY = y+1;
                        break;
                    case '-':
                        prevX = x-1;
                        prevY = y;
                        nextX = x+1;
                        nextY = y;
                        break;
                    case 'L':
                        prevX = x;
                        prevY = y-1;
                        nextX = x+1;
                        nextY = y;
                        break;
                    case 'J':
                        prevX = x;
                        prevY = y-1;
                        nextX = x-1;
                        nextY = y;
                        break;
                    case '7':
                        prevX = x;
                        prevY = y+1;
                        nextX = x-1;
                        nextY = y;
                        break;
                    case 'F':
                        prevX = x;
                        prevY = y+1;
                        nextX = x+1;
                        nextY = y;
                        break;
                    default:
                        prevX = -1;
                        prevY = -1;
                        nextX = -1;
                        nextY = -1;
                        break;
                }
                if (prevX >= 0 && prevY >= 0 && prevX < maxX && prevY < maxY){
                    Pipe prev = pipes[prevY][prevX];
                    pipe.prev = prev;
                    if (prev.type == 'S'){
                        if (prev.next == null) prev.next = pipe;
                        else prev.prev = pipe;
                    }
                }
                if (nextX >= 0 && nextY >= 0 && nextX < maxX && nextY < maxY){
                    Pipe next = pipes[nextY][nextX];
                    pipe.next = next;
                    if (next.type == 'S') {
                        if (next.prev == null) next.prev = pipe;
                        else next.next = pipe;
                    }
                }
            }
        }

    }

    static void printPipes(){
        for (int y=0; y<pipes.length; y++){
            for (int x=0; x<pipes[y].length; x++){
                System.out.print(pipes[y][x]);
            }
            System.out.println();
        }
    }

    static void writePipes(){
        List<Pipe> loop = new ArrayList<>();
        Pipe latest = null;
        Pipe current = start;
        while((current.next != start && current.prev != start) || latest == start){
            Pipe tmp;
            if (current.next != latest) tmp = current.next;
            else if (current.prev != latest) tmp = current.prev;
            else throw new RuntimeException("No valid next or prev");
            loop.add(current);
            latest = current;
            current = tmp;
        }
        loop.add(current);
        char[][] map = new char[pipes.length][pipes[0].length];
        for (int y=0; y<pipes.length; y++){
            for (int x=0; x<pipes[y].length; x++){
                if (loop.contains(pipes[y][x])) map[y][x] = pipes[y][x].type;
                else map[y][x] = '.';
            }
        }
        String[] mapStr = new String[map.length];
        for (int y=0; y<map.length; y++){
            String str = new String(map[y]);
            mapStr[y] = str+"\n";
        }
        InputFile.writeFile(mapStr, "./Tim3/src/input/10Map.txt");
    }

    static void writePipesReachable(){
        List<Pipe> loop = new ArrayList<>();
        Pipe latest = null;
        Pipe current = start;
        while((current.next != start && current.prev != start) || latest == start){
            Pipe tmp;
            if (current.next != latest) tmp = current.next;
            else if (current.prev != latest) tmp = current.prev;
            else throw new RuntimeException("No valid next or prev");
            loop.add(current);
            latest = current;
            current = tmp;
        }
        loop.add(current);
        char[][] map = new char[pipes.length][pipes[0].length];
        for (int y=0; y<pipes.length; y++){
            for (int x=0; x<pipes[y].length; x++){
                if (loop.contains(pipes[y][x])) map[y][x] = pipes[y][x].type;
                else if (!pipes[y][x].reachableFromOutsideLoop) map[y][x] = '0';
                else map[y][x] = '.';
            }
        }
        String[] mapStr = new String[map.length];
        for (int y=0; y<map.length; y++){
            String str = new String(map[y]);
            mapStr[y] = str+"\n";
        }
        InputFile.writeFile(mapStr, "./Tim3/src/input/10Map1.txt");
    }

    static int part1(){
        Pipe latest = null;
        Pipe current = start;
        int steps = 0;
        while((current.next != start && current.prev != start) || latest == start){
            steps++;
            Pipe tmp;
            if (current.next != latest) tmp = current.next;
            else if (current.prev != latest) tmp = current.prev;
            else throw new RuntimeException("No valid next or prev");
            latest = current;
            current = tmp;
        }
        return (steps+1)/2;
    }

    static int part2(){
        List<Pipe> loop = new ArrayList<>();
        Pipe latest = null;
        Pipe current = start;
        while((current.next != start && current.prev != start) || latest == start){
            Pipe tmp;
            if (current.next != latest) tmp = current.next;
            else if (current.prev != latest) tmp = current.prev;
            else throw new RuntimeException("No valid next or prev");
            loop.add(current);
            latest = current;
            current = tmp;
        }
        loop.add(current);
        //loop through rows from start
        for (int y = 0; y < pipes.length; y++) {
            for (int x = 0; x < pipes[y].length; x++) {
                Pipe pipe = pipes[y][x];
                if (x == 0) pipe.reachableFromOutsideLoop = true;
                else if (x-1 >= 0 && pipes[y][x-1].reachableFromOutsideLoop) {
                    if (pipes[y][x-1].type != '|'){
                        pipe.reachableFromOutsideLoop = true;
                    }
                }
            }
        }
        //loop through rows from end
        for (int y = pipes.length-1; y >= 0; y--) {
            for (int x = pipes[y].length-1; x >= 0; x--) {
                Pipe pipe = pipes[y][x];
                if (x == pipes[y].length-1) pipe.reachableFromOutsideLoop = true;
                else if (x+1 < pipes[y].length && pipes[y][x+1].reachableFromOutsideLoop) {
                    if (pipes[y][x+1].type != '|'){
                        pipe.reachableFromOutsideLoop = true;
                    }
                }else if (x-1 >= 0 && pipes[y][x-1].reachableFromOutsideLoop) {
                    if (pipes[y][x-1].type != '|'){
                        pipe.reachableFromOutsideLoop = true;
                    }
                }
            }
        }
        //loop through columns from start
        for (int x = 0; x < pipes[0].length; x++) {
            for (int y = 0; y < pipes.length; y++) {
                Pipe pipe = pipes[y][x];
                if (y == 0) pipe.reachableFromOutsideLoop = true;
                else if (y-1 >= 0 && pipes[y-1][x].reachableFromOutsideLoop) {
                    if (pipes[y-1][x].type != '-'){
                        pipe.reachableFromOutsideLoop = true;
                    }
                }else if (x+1 < pipes[y].length && pipes[y][x+1].reachableFromOutsideLoop) {
                    if (pipes[y][x+1].type != '|'){
                        pipe.reachableFromOutsideLoop = true;
                    }
                }else if (x-1 >= 0 && pipes[y][x-1].reachableFromOutsideLoop) {
                    if (pipes[y][x-1].type != '|'){
                        pipe.reachableFromOutsideLoop = true;
                    }
                }
            }
        }
        //loop through columns from end
        for (int x = pipes[0].length-1; x >= 0; x--) {
            for (int y = pipes.length-1; y >= 0; y--) {
                Pipe pipe = pipes[y][x];
                if (y == pipes.length-1) pipe.reachableFromOutsideLoop = true;
                else if (y+1 < pipes.length && pipes[y+1][x].reachableFromOutsideLoop) {
                    if (pipes[y+1][x].type != '-'){
                        pipe.reachableFromOutsideLoop = true;
                    }
                }
                else if (y-1 >= 0 && pipes[y-1][x].reachableFromOutsideLoop) {
                    if (pipes[y-1][x].type != '-'){
                        pipe.reachableFromOutsideLoop = true;
                    }
                }else if (x+1 < pipes[y].length && pipes[y][x+1].reachableFromOutsideLoop) {
                    if (pipes[y][x+1].type != '|'){
                        pipe.reachableFromOutsideLoop = true;
                    }
                }else if (x-1 >= 0 && pipes[y][x-1].reachableFromOutsideLoop) {
                    if (pipes[y][x-1].type != '|'){
                        pipe.reachableFromOutsideLoop = true;
                    }
                }
                else break;
            }
        }
        // loop through rows from start
        for (int y = 0; y < pipes.length; y++) {
            for (int x = 0; x < pipes[y].length; x++) {
                Pipe pipe = pipes[y][x];
                if (x == 0) pipe.reachableFromOutsideLoop = true;
                else if (y+1 < pipes.length && pipes[y+1][x].reachableFromOutsideLoop) {
                    if (pipes[y+1][x].type != '-'){
                        pipe.reachableFromOutsideLoop = true;
                    }
                }
                else if (y-1 >= 0 && pipes[y-1][x].reachableFromOutsideLoop) {
                    if (pipes[y-1][x].type != '-'){
                        pipe.reachableFromOutsideLoop = true;
                    }
                }else if (x+1 < pipes[y].length && pipes[y][x+1].reachableFromOutsideLoop) {
                    if (pipes[y][x+1].type != '|'){
                        pipe.reachableFromOutsideLoop = true;
                    }
                }else if (x-1 >= 0 && pipes[y][x-1].reachableFromOutsideLoop) {
                    if (pipes[y][x-1].type != '|'){
                        pipe.reachableFromOutsideLoop = true;
                    }
                }
            }
        }
        // loop through rows from end
        for (int y = pipes.length-1; y >= 0; y--) {
            for (int x = pipes[y].length-1; x >= 0; x--) {
                Pipe pipe = pipes[y][x];
                if (x == pipes[y].length-1) pipe.reachableFromOutsideLoop = true;
                else if (y+1 < pipes.length && pipes[y+1][x].reachableFromOutsideLoop) {
                    if (pipes[y+1][x].type != '-'){
                        pipe.reachableFromOutsideLoop = true;
                    }
                }
                else if (y-1 >= 0 && pipes[y-1][x].reachableFromOutsideLoop) {
                    if (pipes[y-1][x].type != '-'){
                        pipe.reachableFromOutsideLoop = true;
                    }
                }else if (x+1 < pipes[y].length && pipes[y][x+1].reachableFromOutsideLoop) {
                    if (pipes[y][x+1].type != '|'){
                        pipe.reachableFromOutsideLoop = true;
                    }
                }else if (x-1 >= 0 && pipes[y][x-1].reachableFromOutsideLoop) {
                    if (pipes[y][x-1].type != '|'){
                        pipe.reachableFromOutsideLoop = true;
                    }
                }
            }
        }
        // loop through columns from start
        for (int x = 0; x < pipes[0].length; x++) {
            for (int y = 0; y < pipes.length; y++) {
                Pipe pipe = pipes[y][x];
                if (y == 0) pipe.reachableFromOutsideLoop = true;
                else if (y+1 < pipes.length && pipes[y+1][x].reachableFromOutsideLoop) {
                    if (pipes[y+1][x].type != '-'){
                        pipe.reachableFromOutsideLoop = true;
                    }
                }
                else if (y-1 >= 0 && pipes[y-1][x].reachableFromOutsideLoop) {
                    if (pipes[y-1][x].type != '-'){
                        pipe.reachableFromOutsideLoop = true;
                    }
                }else if (x+1 < pipes[y].length && pipes[y][x+1].reachableFromOutsideLoop) {
                    if (pipes[y][x+1].type != '|'){
                        pipe.reachableFromOutsideLoop = true;
                    }
                }else if (x-1 >= 0 && pipes[y][x-1].reachableFromOutsideLoop) {
                    if (pipes[y][x-1].type != '|'){
                        pipe.reachableFromOutsideLoop = true;
                    }
                }
            }
        }
        // loop through columns from end
        for (int x = pipes[0].length-1; x >= 0; x--) {
            for (int y = pipes.length-1; y >= 0; y--) {
                Pipe pipe = pipes[y][x];
                if (y == pipes.length-1) pipe.reachableFromOutsideLoop = true;
                else if (y+1 < pipes.length && pipes[y+1][x].reachableFromOutsideLoop) {
                    if (pipes[y+1][x].type != '-'){
                        pipe.reachableFromOutsideLoop = true;
                    }
                }
                else if (y-1 >= 0 && pipes[y-1][x].reachableFromOutsideLoop) {
                    if (pipes[y-1][x].type != '-'){
                        pipe.reachableFromOutsideLoop = true;
                    }
                }else if (x+1 < pipes[y].length && pipes[y][x+1].reachableFromOutsideLoop) {
                    if (pipes[y][x+1].type != '|'){
                        pipe.reachableFromOutsideLoop = true;
                    }
                }else if (x-1 >= 0 && pipes[y][x-1].reachableFromOutsideLoop) {
                    if (pipes[y][x-1].type != '|'){
                        pipe.reachableFromOutsideLoop = true;
                    }
                }
            }
        }

        int c = 0;
        for (Pipe[] row: pipes){
            for (Pipe pipe: row){
                if (!pipe.reachableFromOutsideLoop && !loop.contains(pipe)) c++;
            }
        }
        return c;
    }

    public static void main(String[] args) {
        parseInput();
        //printPipes();
        System.out.println(part1());
        System.out.println(part2());
        writePipesReachable();
    }

}
