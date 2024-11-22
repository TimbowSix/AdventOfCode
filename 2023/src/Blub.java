import java.util.UUID;

public class Blub {
    public static void main(String[] args) {
        UUID uuid = UUID.fromString("00026f67-42f6-43f7-aa69-4263ec126c27");
        //System.out.println(uuid.variant());
        //System.out.println(uuid.version());

        UUID uuid2 = UUID.fromString("9e9a8e34-b034-8c9e-da10-d222e84b7704");
        System.out.println(uuid2.variant());
        System.out.println(uuid2.version());
        //System.out.println(uuid2.clockSequence());
        //System.out.println(uuid2.timestamp());
    }

}
