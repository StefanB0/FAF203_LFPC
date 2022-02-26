module com.lfpc.lfpc_code {
    requires javafx.controls;
    requires javafx.fxml;
    requires org.json;
    requires org.jgrapht.core;
    requires org.jgrapht.io;


    opens com.lfpc.lfpc_code to javafx.fxml;
    exports com.lfpc.lfpc_code;
}